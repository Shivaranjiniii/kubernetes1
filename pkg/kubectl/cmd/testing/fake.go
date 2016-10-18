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

package testing

import (
	"errors"
	"fmt"
	"io"

	"github.com/emicklei/go-restful/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/testapi"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/typed/discovery"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/fake"
	"k8s.io/kubernetes/pkg/kubectl"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/runtime/serializer"
)

type internalType struct {
	Kind       string
	APIVersion string

	Name string
}

type externalType struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`

	Name string `json:"name"`
}

type ExternalType2 struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`

	Name string `json:"name"`
}

func (obj *internalType) GetObjectKind() unversioned.ObjectKind { return obj }
func (obj *internalType) SetGroupVersionKind(gvk unversioned.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}
func (obj *internalType) GroupVersionKind() unversioned.GroupVersionKind {
	return unversioned.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}
func (obj *externalType) GetObjectKind() unversioned.ObjectKind { return obj }
func (obj *externalType) SetGroupVersionKind(gvk unversioned.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}
func (obj *externalType) GroupVersionKind() unversioned.GroupVersionKind {
	return unversioned.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}
func (obj *ExternalType2) GetObjectKind() unversioned.ObjectKind { return obj }
func (obj *ExternalType2) SetGroupVersionKind(gvk unversioned.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}
func (obj *ExternalType2) GroupVersionKind() unversioned.GroupVersionKind {
	return unversioned.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}

func NewInternalType(kind, apiversion, name string) *internalType {
	item := internalType{Kind: kind,
		APIVersion: apiversion,
		Name:       name}
	return &item
}

var versionErr = errors.New("not a version")

func versionErrIfFalse(b bool) error {
	if b {
		return nil
	}
	return versionErr
}

var validVersion = registered.GroupOrDie(api.GroupName).GroupVersion.Version
var internalGV = unversioned.GroupVersion{Group: "apitest", Version: runtime.APIVersionInternal}
var UnlikelyGV = unversioned.GroupVersion{Group: "apitest", Version: "unlikelyversion"}
var validVersionGV = unversioned.GroupVersion{Group: "apitest", Version: validVersion}

func newExternalScheme() (*runtime.Scheme, meta.RESTMapper, runtime.Codec) {
	scheme := runtime.NewScheme()
	scheme.AddKnownTypeWithName(internalGV.WithKind("Type"), &internalType{})
	scheme.AddKnownTypeWithName(UnlikelyGV.WithKind("Type"), &externalType{})
	//This tests that kubectl will not confuse the external scheme with the internal scheme, even when they accidentally have versions of the same name.
	scheme.AddKnownTypeWithName(validVersionGV.WithKind("Type"), &ExternalType2{})

	codecs := serializer.NewCodecFactory(scheme)
	codec := codecs.LegacyCodec(UnlikelyGV)
	mapper := meta.NewDefaultRESTMapper([]unversioned.GroupVersion{UnlikelyGV, validVersionGV}, func(version unversioned.GroupVersion) (*meta.VersionInterfaces, error) {
		return &meta.VersionInterfaces{
			ObjectConvertor:  scheme,
			MetadataAccessor: meta.NewAccessor(),
		}, versionErrIfFalse(version == validVersionGV || version == UnlikelyGV)
	})
	for _, gv := range []unversioned.GroupVersion{UnlikelyGV, validVersionGV} {
		for kind := range scheme.KnownTypes(gv) {
			gvk := gv.WithKind(kind)

			scope := meta.RESTScopeNamespace
			mapper.Add(gvk, scope)
		}
	}

	return scheme, mapper, codec
}

// TODO move to fake.go

type testFactory struct {
	Mapper       meta.RESTMapper
	Typer        runtime.ObjectTyper
	Client       kubectl.RESTClient
	Describer    kubectl.Describer
	Printer      kubectl.ResourcePrinter
	Validator    validation.Schema
	Namespace    string
	ClientConfig *restclient.Config
	Err          error
}

type fakeFactory struct {
	tf    *testFactory
	Codec runtime.Codec
}

func NewTestFactory() (cmdutil.Factory, *testFactory, runtime.Codec, runtime.NegotiatedSerializer) {
	scheme, mapper, codec := newExternalScheme()
	t := &testFactory{
		Validator: validation.NullSchema{},
		Mapper:    mapper,
		Typer:     scheme,
	}
	negotiatedSerializer := serializer.NegotiatedSerializerWrapper(
		runtime.SerializerInfo{Serializer: codec},
		runtime.StreamSerializerInfo{})
	return &fakeFactory{
		tf:    t,
		Codec: codec,
	}, t, codec, negotiatedSerializer
}

func (f *fakeFactory) FlagSet() *pflag.FlagSet {
	return nil
}

func (f *fakeFactory) Object() (meta.RESTMapper, runtime.ObjectTyper) {
	priorityRESTMapper := meta.PriorityRESTMapper{
		Delegate: f.tf.Mapper,
		ResourcePriority: []unversioned.GroupVersionResource{
			{Group: meta.AnyGroup, Version: "v1", Resource: meta.AnyResource},
		},
		KindPriority: []unversioned.GroupVersionKind{
			{Group: meta.AnyGroup, Version: "v1", Kind: meta.AnyKind},
		},
	}
	return priorityRESTMapper, f.tf.Typer
}

func (f *fakeFactory) UnstructuredObject() (meta.RESTMapper, runtime.ObjectTyper, error) {
	return nil, nil, nil
}

func (f *fakeFactory) Decoder(bool) runtime.Decoder {
	return f.Codec
}

func (f *fakeFactory) JSONEncoder() runtime.Encoder {
	return f.Codec
}

func (f *fakeFactory) RESTClient() (*restclient.RESTClient, error) {
	return nil, nil
}

func (f *fakeFactory) ClientSet() (*internalclientset.Clientset, error) {
	return nil, nil
}

func (f *fakeFactory) ClientConfig() (*restclient.Config, error) {
	return f.tf.ClientConfig, f.tf.Err
}

func (f *fakeFactory) ClientForMapping(*meta.RESTMapping) (resource.RESTClient, error) {
	return f.tf.Client, f.tf.Err
}

func (f *fakeFactory) UnstructuredClientForMapping(*meta.RESTMapping) (resource.RESTClient, error) {
	return nil, nil
}

func (f *fakeFactory) Describer(*meta.RESTMapping) (kubectl.Describer, error) {
	return f.tf.Describer, f.tf.Err
}

func (f *fakeFactory) Printer(mapping *meta.RESTMapping, options kubectl.PrintOptions) (kubectl.ResourcePrinter, error) {
	return f.tf.Printer, f.tf.Err
}

func (f *fakeFactory) Scaler(*meta.RESTMapping) (kubectl.Scaler, error) {
	return nil, nil
}

func (f *fakeFactory) Reaper(*meta.RESTMapping) (kubectl.Reaper, error) {
	return nil, nil
}

func (f *fakeFactory) HistoryViewer(*meta.RESTMapping) (kubectl.HistoryViewer, error) {
	return nil, nil
}

func (f *fakeFactory) Rollbacker(*meta.RESTMapping) (kubectl.Rollbacker, error) {
	return nil, nil
}

func (f *fakeFactory) StatusViewer(*meta.RESTMapping) (kubectl.StatusViewer, error) {
	return nil, nil
}

func (f *fakeFactory) MapBasedSelectorForObject(runtime.Object) (string, error) {
	return "", nil
}

func (f *fakeFactory) PortsForObject(runtime.Object) ([]string, error) {
	return nil, nil
}

func (f *fakeFactory) ProtocolsForObject(runtime.Object) (map[string]string, error) {
	return nil, nil
}

func (f *fakeFactory) LabelsForObject(runtime.Object) (map[string]string, error) {
	return nil, nil
}

func (f *fakeFactory) LogsForObject(object, options runtime.Object) (*restclient.Request, error) {
	return nil, nil
}

func (f *fakeFactory) PauseObject(runtime.Object) (bool, error) {
	return false, nil
}

func (f *fakeFactory) ResumeObject(runtime.Object) (bool, error) {
	return false, nil
}

func (f *fakeFactory) Validator(validate bool, cacheDir string) (validation.Schema, error) {
	return f.tf.Validator, f.tf.Err
}

func (f *fakeFactory) SwaggerSchema(unversioned.GroupVersionKind) (*swagger.ApiDeclaration, error) {
	return nil, nil
}

func (f *fakeFactory) DefaultNamespace() (string, bool, error) {
	return f.tf.Namespace, false, f.tf.Err
}

func (f *fakeFactory) Generators(string) map[string]kubectl.Generator {
	return nil
}

func (f *fakeFactory) CanBeExposed(unversioned.GroupKind) error {
	return nil
}

func (f *fakeFactory) CanBeAutoscaled(unversioned.GroupKind) error {
	return nil
}

func (f *fakeFactory) AttachablePodForObject(ob runtime.Object) (*api.Pod, error) {
	return nil, nil
}

func (f *fakeFactory) UpdatePodSpecForObject(obj runtime.Object, fn func(*api.PodSpec) error) (bool, error) {
	return false, nil
}

func (f *fakeFactory) EditorEnvs() []string {
	return nil
}

func (f *fakeFactory) PrintObjectSpecificMessage(obj runtime.Object, out io.Writer) {
}

func (f *fakeFactory) Command() string {
	return ""
}

func (f *fakeFactory) BindFlags(flags *pflag.FlagSet) {
}

func (f *fakeFactory) BindExternalFlags(flags *pflag.FlagSet) {
}

func (f *fakeFactory) PrintObject(cmd *cobra.Command, mapper meta.RESTMapper, obj runtime.Object, out io.Writer) error {
	return nil
}

func (f *fakeFactory) PrinterForMapping(cmd *cobra.Command, mapping *meta.RESTMapping, withNamespace bool) (kubectl.ResourcePrinter, error) {
	return f.tf.Printer, f.tf.Err
}

func (f *fakeFactory) NewBuilder() *resource.Builder {
	return nil
}

func (f *fakeFactory) DefaultResourceFilterOptions(cmd *cobra.Command, withNamespace bool) *kubectl.PrintOptions {
	return &kubectl.PrintOptions{}
}

func (f *fakeFactory) DefaultResourceFilterFunc() kubectl.Filters {
	return nil
}

type fakeMixedFactory struct {
	cmdutil.Factory
	tf        *testFactory
	apiClient resource.RESTClient
}

func (f *fakeMixedFactory) Object() (meta.RESTMapper, runtime.ObjectTyper) {
	var multiRESTMapper meta.MultiRESTMapper
	multiRESTMapper = append(multiRESTMapper, f.tf.Mapper)
	multiRESTMapper = append(multiRESTMapper, testapi.Default.RESTMapper())
	priorityRESTMapper := meta.PriorityRESTMapper{
		Delegate: multiRESTMapper,
		ResourcePriority: []unversioned.GroupVersionResource{
			{Group: meta.AnyGroup, Version: "v1", Resource: meta.AnyResource},
		},
		KindPriority: []unversioned.GroupVersionKind{
			{Group: meta.AnyGroup, Version: "v1", Kind: meta.AnyKind},
		},
	}
	return priorityRESTMapper, runtime.MultiObjectTyper{f.tf.Typer, api.Scheme}
}

func (f *fakeMixedFactory) ClientForMapping(m *meta.RESTMapping) (resource.RESTClient, error) {
	if m.ObjectConvertor == api.Scheme {
		return f.apiClient, f.tf.Err
	}
	return f.tf.Client, f.tf.Err
}

func NewMixedFactory(apiClient resource.RESTClient) (cmdutil.Factory, *testFactory, runtime.Codec) {
	f, t, c, _ := NewTestFactory()
	return &fakeMixedFactory{
		Factory:   f,
		tf:        t,
		apiClient: apiClient,
	}, t, c
}

type fakeAPIFactory struct {
	cmdutil.Factory
	tf *testFactory
}

func (f *fakeAPIFactory) Object() (meta.RESTMapper, runtime.ObjectTyper) {
	return testapi.Default.RESTMapper(), api.Scheme
}

func (f *fakeAPIFactory) UnstructuredObject() (meta.RESTMapper, runtime.ObjectTyper, error) {
	groupResources := testDynamicResources()
	mapper := discovery.NewRESTMapper(groupResources, meta.InterfacesForUnstructured)
	typer := discovery.NewUnstructuredObjectTyper(groupResources)

	return cmdutil.NewShortcutExpander(mapper, nil), typer, nil
}

func (f *fakeAPIFactory) Decoder(bool) runtime.Decoder {
	return testapi.Default.Codec()
}

func (f *fakeAPIFactory) JSONEncoder() runtime.Encoder {
	return testapi.Default.Codec()
}

func (f *fakeAPIFactory) ClientSet() (*internalclientset.Clientset, error) {
	// Swap out the HTTP client out of the client with the fake's version.
	fakeClient := f.tf.Client.(*fake.RESTClient)
	restClient, err := restclient.RESTClientFor(f.tf.ClientConfig)
	if err != nil {
		panic(err)
	}
	restClient.Client = fakeClient.Client
	return internalclientset.New(restClient), f.tf.Err
}

func (f *fakeAPIFactory) RESTClient() (*restclient.RESTClient, error) {
	// Swap out the HTTP client out of the client with the fake's version.
	fakeClient := f.tf.Client.(*fake.RESTClient)
	restClient, err := restclient.RESTClientFor(f.tf.ClientConfig)
	if err != nil {
		panic(err)
	}
	restClient.Client = fakeClient.Client
	return restClient, f.tf.Err
}

func (f *fakeAPIFactory) ClientConfig() (*restclient.Config, error) {
	return f.tf.ClientConfig, f.tf.Err
}

func (f *fakeAPIFactory) ClientForMapping(*meta.RESTMapping) (resource.RESTClient, error) {
	return f.tf.Client, f.tf.Err
}

func (f *fakeAPIFactory) UnstructuredClientForMapping(*meta.RESTMapping) (resource.RESTClient, error) {
	return f.tf.Client, f.tf.Err
}

func (f *fakeAPIFactory) Describer(*meta.RESTMapping) (kubectl.Describer, error) {
	return f.tf.Describer, f.tf.Err
}

func (f *fakeAPIFactory) Printer(mapping *meta.RESTMapping, options kubectl.PrintOptions) (kubectl.ResourcePrinter, error) {
	return f.tf.Printer, f.tf.Err
}

func (f *fakeAPIFactory) LogsForObject(object, options runtime.Object) (*restclient.Request, error) {
	fakeClient := f.tf.Client.(*fake.RESTClient)
	c := client.NewOrDie(f.tf.ClientConfig)
	c.Client = fakeClient.Client

	switch t := object.(type) {
	case *api.Pod:
		opts, ok := options.(*api.PodLogOptions)
		if !ok {
			return nil, errors.New("provided options object is not a PodLogOptions")
		}
		return c.Pods(f.tf.Namespace).GetLogs(t.Name, opts), nil
	default:
		fqKinds, _, err := api.Scheme.ObjectKinds(object)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("cannot get the logs from %v", fqKinds[0])
	}
}

func (f *fakeAPIFactory) Validator(validate bool, cacheDir string) (validation.Schema, error) {
	return f.tf.Validator, f.tf.Err
}

func (f *fakeAPIFactory) DefaultNamespace() (string, bool, error) {
	return f.tf.Namespace, false, f.tf.Err
}

func (f *fakeAPIFactory) Generators(cmdName string) map[string]kubectl.Generator {
	return cmdutil.DefaultGenerators(cmdName)
}

func (f *fakeAPIFactory) PrintObject(cmd *cobra.Command, mapper meta.RESTMapper, obj runtime.Object, out io.Writer) error {
	gvks, _, err := api.Scheme.ObjectKinds(obj)
	if err != nil {
		return err
	}

	mapping, err := mapper.RESTMapping(gvks[0].GroupKind())
	if err != nil {
		return err
	}

	printer, err := f.PrinterForMapping(cmd, mapping, false)
	if err != nil {
		return err
	}
	return printer.PrintObj(obj, out)
}

func (f *fakeAPIFactory) PrinterForMapping(cmd *cobra.Command, mapping *meta.RESTMapping, withNamespace bool) (kubectl.ResourcePrinter, error) {
	return f.tf.Printer, f.tf.Err
}

func (f *fakeAPIFactory) NewBuilder() *resource.Builder {
	mapper, typer := f.Object()

	return resource.NewBuilder(mapper, typer, resource.ClientMapperFunc(f.ClientForMapping), f.Decoder(true))
}

func NewAPIFactory() (cmdutil.Factory, *testFactory, runtime.Codec, runtime.NegotiatedSerializer) {
	t := &testFactory{
		Validator: validation.NullSchema{},
	}
	rf := cmdutil.NewFactory(nil)
	return &fakeAPIFactory{
		Factory: rf,
		tf:      t,
	}, t, testapi.Default.Codec(), testapi.Default.NegotiatedSerializer()
}

func testDynamicResources() []*discovery.APIGroupResources {
	return []*discovery.APIGroupResources{
		{
			Group: unversioned.APIGroup{
				Versions: []unversioned.GroupVersionForDiscovery{
					{Version: "v1"},
				},
				PreferredVersion: unversioned.GroupVersionForDiscovery{Version: "v1"},
			},
			VersionedResources: map[string][]unversioned.APIResource{
				"v1": {
					{Name: "pods", Namespaced: true, Kind: "Pod"},
					{Name: "services", Namespaced: true, Kind: "Service"},
					{Name: "replicationcontrollers", Namespaced: true, Kind: "ReplicationController"},
				},
			},
		},
	}
}
