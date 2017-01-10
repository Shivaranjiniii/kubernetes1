/*
Copyright 2017 The Kubernetes Authors.

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

package kubectl

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/rbac"
)

// RoleGeneratorV1 supports stable generation of a role.
type RoleGeneratorV1 struct {
	// Name of role (required)
	Name string
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds contained in this rule (required)
	Verbs []string
	// Resources is a list of resources this rule applies to (required)
	Resources []string
	// APIGroups is the name of the APIGroup that contains the resources (optional)
	APIGroups []string
	// ResourceNames is an optional white list of names that the rule applies to (optional)
	ResourceNames []string
	// Mapper is used to validate resources.
	Mapper meta.RESTMapper
}

// Ensure it supports the generator pattern that uses parameter injection.
var _ Generator = &RoleGeneratorV1{}

// Ensure it supports the generator pattern that uses parameters specified during construction.
var _ StructuredGenerator = &RoleGeneratorV1{}

// Valid verb list for validation.
var validVerbs []string = []string{"*", "get", "post", "put", "delete", "list", "create", "update", "patch", "watch", "proxy", "redirect", "deletecollection"}

// Generate returns a role using the specified parameters.
func (s *RoleGeneratorV1) Generate(genericParams map[string]interface{}) (runtime.Object, error) {
	err := ValidateParams(s.ParamNames(), genericParams)
	if err != nil {
		return nil, err
	}

	delegate := &RoleGeneratorV1{Mapper: s.Mapper}

	name, isString := genericParams["name"].(string)
	if !isString {
		return nil, fmt.Errorf("expected string, found: %v", name)
	}
	delegate.Name = name

	verbs, isArray := genericParams["verb"].([]string)
	if !isArray {
		return nil, fmt.Errorf("expected []string, found: %v", verbs)
	}
	delegate.Verbs = verbs

	resources, isArray := genericParams["resource"].([]string)
	if !isArray {
		return nil, fmt.Errorf("expected []string, found: %v", resources)
	}
	delegate.Resources = resources

	apiGroupStrings, found := genericParams["api-group"]
	if found {
		apiGroups, isArray := apiGroupStrings.([]string)
		if !isArray {
			return nil, fmt.Errorf("expected []string, found: %v", apiGroupStrings)
		}
		delegate.APIGroups = apiGroups
	}

	resourceNameStrings, found := genericParams["resource-name"]
	if found {
		resourceNames, isArray := resourceNameStrings.([]string)
		if !isArray {
			return nil, fmt.Errorf("expected []string, found :%v", resourceNameStrings)
		}
		delegate.ResourceNames = resourceNames
	}

	return delegate.StructuredGenerate()
}

// ParamNames returns the set of supported input parameters when using the parameter injection generator pattern.
func (s *RoleGeneratorV1) ParamNames() []GeneratorParam {
	return []GeneratorParam{
		{"name", true},
		{"verb", true},
		{"resource", true},
		{"api-group", false},
		{"resource-name", false},
	}
}

// StructuredGenerate outputs a role object using the configured fields.
func (s *RoleGeneratorV1) StructuredGenerate() (runtime.Object, error) {
	if err := s.parseAndValidateParams(); err != nil {
		return nil, err
	}

	role := &rbac.Role{}
	rule := rbac.PolicyRule{}

	role.Name = s.Name
	rule.Verbs = s.Verbs
	rule.Resources = s.Resources
	rule.APIGroups = s.APIGroups
	rule.ResourceNames = s.ResourceNames

	// At present, we only allow creating one single rule by using 'kubectl create role' command.
	role.Rules = []rbac.PolicyRule{rule}

	return role, nil
}

// parseAndValidateParams validates required fields are set to support structured generation.
func (s *RoleGeneratorV1) parseAndValidateParams() error {
	if len(s.Name) == 0 {
		return fmt.Errorf("name must be specified")
	}

	err := s.parseAndValidateVerbs()
	if err != nil {
		return err
	}

	err = s.parseAndValidateResources()
	if err != nil {
		return err
	}

	err = s.parseAndValidateAPIGroups()
	if err != nil {
		return err
	}

	return s.parseAndValidateResourceNames()
}

// parseAndValidateVerbs parses and validates verbs.
func (s *RoleGeneratorV1) parseAndValidateVerbs() error {
	// support specify multiple verbs together
	// e.g. --verb=get,watch,list
	verbs := []string{}
	for _, v := range s.Verbs {
		verbs = mergeArrays(verbs, strings.Split(v, ","))
	}

	if len(verbs) == 0 {
		return fmt.Errorf("at least one verb must be specified")
	}

	containsVerbAll := false
	for _, v := range verbs {
		if !arrayContains(validVerbs, v) {
			return fmt.Errorf("invalid verb: '%s'", v)
		}
		if v == "*" {
			containsVerbAll = true
		}
	}

	// VerbAll respresents all kinds of verbs.
	if containsVerbAll {
		s.Verbs = []string{"*"}
	} else {
		s.Verbs = verbs
	}
	return nil
}

// parseAndValidateResources parses and validates resources.
func (s *RoleGeneratorV1) parseAndValidateResources() error {
	// support specify multiple resources together
	// e.g. --resource=pods,deployments
	candidateResources := []string{}
	for _, r := range s.Resources {
		candidateResources = mergeArrays(candidateResources, strings.Split(r, ","))
	}

	if len(candidateResources) == 0 {
		return fmt.Errorf("at least one resource must be specified")
	}

	resources := []string{}
	for _, r := range candidateResources {
		// support resource.group pattern
		index := strings.Index(r, ".")
		resource, apiGroup := "", ""

		// No API Group specified, use "" as core API Group
		if index == -1 {
			resource = r
		} else {
			resource, apiGroup = r[0:index], r[index+1:]
		}

		if !arrayContains(s.APIGroups, apiGroup) {
			s.APIGroups = append(s.APIGroups, apiGroup)
		}
		if !arrayContains(resources, resource) {
			resources = append(resources, resource)
		}
	}

	if s.Mapper == nil {
		return fmt.Errorf("no REST Mapper found to validate resources")
	}

	containsResourceAll := false
	for _, r := range resources {
		// ResourceAll respresents all kinds of resources.
		if r == "*" {
			containsResourceAll = true
			continue
		}
		resource, err := s.Mapper.ResourceFor(schema.GroupVersionResource{Resource: r})
		if err != nil {
			return err
		}
		if resource.Resource != r {
			return fmt.Errorf("invalid resource '%s', do you mean '%s'?", r, resource.Resource)
		}
	}

	if containsResourceAll {
		s.Resources = []string{"*"}
	} else {
		s.Resources = resources
	}
	return nil
}

// parseAndValidateAPIGroups parses and validates api groups.
func (s *RoleGeneratorV1) parseAndValidateAPIGroups() error {
	// support specify multiple api groups together
	// e.g. --api-group=extensions,metrics
	groups := []string{}
	for _, g := range s.APIGroups {
		groups = mergeArrays(groups, strings.Split(g, ","))
	}

	if len(groups) == 0 {
		return fmt.Errorf("at least one api group must be specified")
	}

	for _, g := range groups {
		// APIGroupAll represents all the api groups.
		if g == "*" {
			s.APIGroups = []string{"*"}
			return nil
		}
	}

	s.APIGroups = groups
	return nil
}

// parseAndValidateResourceNames parses and validates resource names.
func (s *RoleGeneratorV1) parseAndValidateResourceNames() error {
	// support specify multiple resource names together
	// e.g. --resource-name=foo,boo
	names := []string{}
	for _, n := range s.ResourceNames {
		names = mergeArrays(names, strings.Split(n, ","))
	}

	if len(names) > 0 && len(s.Resources) > 1 {
		return fmt.Errorf("resource name(s) can not be applied to multiple resources")
	}

	if len(names) > 0 {
		s.ResourceNames = names
	}
	return nil
}

// mergeArrays merges two string arrays with no duplicate element.
func mergeArrays(a []string, b []string) []string {
	for _, v := range b {
		if !arrayContains(a, v) {
			a = append(a, v)
		}
	}
	return a
}
