/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsResourceType Document type to apply an assertion against.
type SyntheticsResourceType string

// List of SyntheticsResourceType
const (
	SYNTHETICSRESOURCETYPE_DOCUMENT   SyntheticsResourceType = "document"
	SYNTHETICSRESOURCETYPE_STYLESHEET SyntheticsResourceType = "stylesheet"
	SYNTHETICSRESOURCETYPE_FETCH      SyntheticsResourceType = "fetch"
	SYNTHETICSRESOURCETYPE_IMAGE      SyntheticsResourceType = "image"
	SYNTHETICSRESOURCETYPE_SCRIPT     SyntheticsResourceType = "script"
	SYNTHETICSRESOURCETYPE_XHR        SyntheticsResourceType = "xhr"
	SYNTHETICSRESOURCETYPE_OTHER      SyntheticsResourceType = "other"
)

// Ptr returns reference to SyntheticsResourceType value
func (v SyntheticsResourceType) Ptr() *SyntheticsResourceType {
	return &v
}

type NullableSyntheticsResourceType struct {
	value *SyntheticsResourceType
	isSet bool
}

func (v NullableSyntheticsResourceType) Get() *SyntheticsResourceType {
	return v.value
}

func (v *NullableSyntheticsResourceType) Set(val *SyntheticsResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsResourceType(val *SyntheticsResourceType) *NullableSyntheticsResourceType {
	return &NullableSyntheticsResourceType{value: val, isSet: true}
}

func (v NullableSyntheticsResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
