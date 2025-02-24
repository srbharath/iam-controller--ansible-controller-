// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package role

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	compareTags(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AssumeRolePolicyDocument, b.ko.Spec.AssumeRolePolicyDocument) {
		delta.Add("Spec.AssumeRolePolicyDocument", a.ko.Spec.AssumeRolePolicyDocument, b.ko.Spec.AssumeRolePolicyDocument)
	} else if a.ko.Spec.AssumeRolePolicyDocument != nil && b.ko.Spec.AssumeRolePolicyDocument != nil {
		if *a.ko.Spec.AssumeRolePolicyDocument != *b.ko.Spec.AssumeRolePolicyDocument {
			delta.Add("Spec.AssumeRolePolicyDocument", a.ko.Spec.AssumeRolePolicyDocument, b.ko.Spec.AssumeRolePolicyDocument)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaxSessionDuration, b.ko.Spec.MaxSessionDuration) {
		delta.Add("Spec.MaxSessionDuration", a.ko.Spec.MaxSessionDuration, b.ko.Spec.MaxSessionDuration)
	} else if a.ko.Spec.MaxSessionDuration != nil && b.ko.Spec.MaxSessionDuration != nil {
		if *a.ko.Spec.MaxSessionDuration != *b.ko.Spec.MaxSessionDuration {
			delta.Add("Spec.MaxSessionDuration", a.ko.Spec.MaxSessionDuration, b.ko.Spec.MaxSessionDuration)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Path, b.ko.Spec.Path) {
		delta.Add("Spec.Path", a.ko.Spec.Path, b.ko.Spec.Path)
	} else if a.ko.Spec.Path != nil && b.ko.Spec.Path != nil {
		if *a.ko.Spec.Path != *b.ko.Spec.Path {
			delta.Add("Spec.Path", a.ko.Spec.Path, b.ko.Spec.Path)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PermissionsBoundary, b.ko.Spec.PermissionsBoundary) {
		delta.Add("Spec.PermissionsBoundary", a.ko.Spec.PermissionsBoundary, b.ko.Spec.PermissionsBoundary)
	} else if a.ko.Spec.PermissionsBoundary != nil && b.ko.Spec.PermissionsBoundary != nil {
		if *a.ko.Spec.PermissionsBoundary != *b.ko.Spec.PermissionsBoundary {
			delta.Add("Spec.PermissionsBoundary", a.ko.Spec.PermissionsBoundary, b.ko.Spec.PermissionsBoundary)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.Policies, b.ko.Spec.Policies) {
		delta.Add("Spec.Policies", a.ko.Spec.Policies, b.ko.Spec.Policies)
	}

	return delta
}
