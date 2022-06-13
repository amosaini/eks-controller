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

package fargate_profile

import (
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"

	svcapitypes "github.com/aws-controllers-k8s/eks-controller/apis/v1alpha1"
)

var (
	_ = svcapitypes.FargateProfile{}
	_ = acktags.NewTags()
)

// ToACKTags converts the tags parameter into 'acktags.Tags' shape.
// This method helps in creating the hub(acktags.Tags) for merging
// default controller tags with existing resource tags.
func ToACKTags(tags map[string]*string) acktags.Tags {
	result := acktags.NewTags()
	if tags == nil || len(tags) == 0 {
		return result
	}

	for k, v := range tags {
		if v == nil {
			result[k] = ""
		} else {
			result[k] = *v
		}
	}

	return result
}

// FromACKTags converts the tags parameter into map[string]*string shape.
// This method helps in setting the tags back inside AWSResource after merging
// default controller tags with existing resource tags.
func FromACKTags(tags acktags.Tags) map[string]*string {
	result := map[string]*string{}
	for k, v := range tags {
		vCopy := v
		result[k] = &vCopy
	}
	return result
}
