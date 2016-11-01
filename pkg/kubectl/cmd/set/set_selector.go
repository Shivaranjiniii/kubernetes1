/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package set

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
	"k8s.io/kubernetes/pkg/runtime"
)

// SelectorOptions is the start of the data required to perform the operation.  As new fields are added, add them here instead of
// referencing the cmd.Flags()
type SelectorOptions struct {
	fnOptions resource.FilenameOptions

	local       bool
	dryrun      bool
	all         bool
	record      bool
	changeCause string

	resources []string
	selector  *unversioned.LabelSelector

	out              io.Writer
	PrintObject      func(obj runtime.Object) error
	ClientForMapping func(mapping *meta.RESTMapping) (resource.RESTClient, error)

	builder *resource.Builder
	mapper  meta.RESTMapper
	encoder runtime.Encoder
}

var (
	selectorLong = templates.LongDesc(`
		Update the selectors on a resource.

		A selector must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to %[1]d characters.
		If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used.
                Note: currently selectors can only be set on Service objects.`)
	selectorExample = templates.Examples(`
                # set the labels and selector before creating a deployment/service pair.
                kubectl create service cluster-ip -o yaml --dry-run | kubectl set selector --local -f - 'environment=qa' -o yaml | kubectl create -f - 
                kubectl create deployment my-dep -o yaml --dry-run | kubectl label --local -f - environment=qa -o yaml | kubectl create -f -`)
)

// NewCmdSelector is the "set selector" command.
func NewCmdSelector(f cmdutil.Factory, out io.Writer) *cobra.Command {
	options := &SelectorOptions{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "selector (-f FILENAME | TYPE NAME) EXPRESSIONS [--resource-version=version]",
		Short:   "Update the selectors on a resource",
		Long:    fmt.Sprintf(selectorLong),
		Example: selectorExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(options.Complete(f, cmd, args, out))
			cmdutil.CheckErr(options.Validate())
			cmdutil.CheckErr(options.RunSelector())
		},
	}
	cmdutil.AddPrinterFlags(cmd)
	cmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	cmd.Flags().Bool("local", false, "If true, set selector will NOT contact api-server but run locally.")
	cmd.Flags().String("resource-version", "", "If non-empty, the selectors update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.")
	usage := "the resource to update the selectors"
	cmdutil.AddFilenameOptionFlags(cmd, &options.fnOptions, usage)
	cmdutil.AddDryRunFlag(cmd)
	cmdutil.AddRecordFlag(cmd)
	cmdutil.AddInclude3rdPartyFlags(cmd)

	return cmd
}

// Complete assigns the SelectorOptions from args.
func (o *SelectorOptions) Complete(f cmdutil.Factory, cmd *cobra.Command, args []string, out io.Writer) error {
	o.local = cmdutil.GetFlagBool(cmd, "local")
	o.all = cmdutil.GetFlagBool(cmd, "all")
	o.record = cmdutil.GetRecordFlag(cmd)
	o.dryrun = cmdutil.GetDryRunFlag(cmd)

	cmdNamespace, enforceNamespace, err := f.DefaultNamespace()
	if err != nil {
		return err
	}

	o.changeCause = f.Command()
	mapper, _ := f.Object()
	o.mapper = mapper
	o.encoder = f.JSONEncoder()

	o.builder = f.NewBuilder().
		ContinueOnError().
		NamespaceParam(cmdNamespace).DefaultNamespace().
		FilenameParam(enforceNamespace, &o.fnOptions).
		Flatten()

	o.PrintObject = func(obj runtime.Object) error {
		return f.PrintObject(cmd, mapper, obj, o.out)
	}
	o.ClientForMapping = func(mapping *meta.RESTMapping) (resource.RESTClient, error) {
		return f.ClientForMapping(mapping)
	}

	o.resources, o.selector, err = getResourcesAndSelector(args)
	return err
}

// Validate basic inputs
func (o *SelectorOptions) Validate() error {
	if len(o.resources) < 1 && cmdutil.IsFilenameEmpty(o.fnOptions.Filenames) {
		return fmt.Errorf("one or more resources must be specified as <resource> <name> or <resource>/<name>")
	}
	return nil
}

// RunSelector executes the command.
func (o *SelectorOptions) RunSelector() error {
	if !o.local {
		o.builder = o.builder.ResourceTypeOrNameArgs(o.all, o.resources...).
			Latest()
	}
	r := o.builder.Do()
	err := r.Err()
	if err != nil {
		return err
	}

	err = r.Visit(func(info *resource.Info, err error) error {
		patch := &Patch{Info: info}
		CalculatePatch(patch, o.encoder, func(info *resource.Info) ([]byte, error) {
			selectErr := updateSelectorForObject(info.Object, *o.selector)

			if selectErr == nil {
				return runtime.Encode(o.encoder, info.Object)
			}
			return nil, selectErr
		})

		if patch.Err != nil {
			return patch.Err
		}
		if o.local || o.dryrun {
			fmt.Fprintln(o.out, "running in local/dry-run mode...")
			o.PrintObject(info.Object)
			return nil
		}

		patched, err := resource.NewHelper(info.Client, info.Mapping).Patch(info.Namespace, info.Name, api.StrategicMergePatchType, patch.Patch)
		if err != nil {
			return err
		}

		// record this change (for rollout history)
		if o.record || cmdutil.ContainsChangeCause(info) {
			if err := cmdutil.RecordChangeCause(patched, o.changeCause); err == nil {
				if patched, err = resource.NewHelper(info.Client, info.Mapping).Replace(info.Namespace, info.Name, false, patched); err != nil {
					return fmt.Errorf("changes to %s/%s can't be recorded: %v\n", info.Mapping.Resource, info.Name, err)
				}
			}
		}

		info.Refresh(patched, true)
		cmdutil.PrintSuccess(o.mapper, false, o.out, info.Mapping.Resource, info.Name, o.dryrun, "selector updated")
		return nil
	})

	return err
}

func updateSelectorForObject(obj runtime.Object, selector unversioned.LabelSelector) error {
	copyOldSelector := func() (map[string]string, error) {
		if len(selector.MatchExpressions) > 0 {
			return nil, fmt.Errorf("match expression %v not supported on this object", selector.MatchExpressions)
		}
		dst := make(map[string]string)
		for label, value := range selector.MatchLabels {
			dst[label] = value
		}
		return dst, nil
	}
	var err error
	switch t := obj.(type) {
	case *api.Service:
		t.Spec.Selector, err = copyOldSelector()
	case *api.PersistentVolumeClaim:
		err = fmt.Errorf("the object %v does not support updates to the selector", t)
	case *extensions.Deployment:
		err = fmt.Errorf("the object %v does not support updates to the selector", t)
	case *extensions.DaemonSet:
		err = fmt.Errorf("the object %v does not support updates to the selector", t)
	case *extensions.ReplicaSet:
		err = fmt.Errorf("the object %v does not support updates to the selector", t)
	case *batch.Job:
		err = fmt.Errorf("the object %v does not support updates to the selector", t)
	case *api.ReplicationController:
		err = fmt.Errorf("the object %v does not support updates to the selector", t)
	default:
		err = fmt.Errorf("the object %v does not have a selector", t)
	}
	return err
}

// getResourcesAndSelector retrieves resources and the selector expression from the given args (assuming selectors the last arg)
func getResourcesAndSelector(args []string) (resources []string, selector *unversioned.LabelSelector, err error) {
	if len(args) > 1 {
		resources = args[:len(args)-1]
	}
	selector, err = unversioned.ParseToLabelSelector(args[len(args)-1])
	return
}
