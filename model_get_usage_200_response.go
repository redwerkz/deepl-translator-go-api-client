/*
DeepL API

The DeepL API provides programmatic access to DeepL’s machine translation technology.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetUsage200Response struct for GetUsage200Response
type GetUsage200Response struct {
	// Characters translated so far in the current billing period.
	CharacterCount *int64 `json:"character_count,omitempty"`
	// Current maximum number of characters that can be translated per billing period.
	CharacterLimit *int64 `json:"character_limit,omitempty"`
	// Documents translated so far in the current billing period.
	DocumentLimit *int64 `json:"document_limit,omitempty"`
	// Current maximum number of documents that can be translated per billing period.
	DocumentCount *int64 `json:"document_count,omitempty"`
	// Documents translated by all users in the team so far in the current billing period.
	TeamDocumentLimit *int64 `json:"team_document_limit,omitempty"`
	// Current maximum number of documents that can be translated by the team per billing period.
	TeamDocumentCount *int64 `json:"team_document_count,omitempty"`
}

// NewGetUsage200Response instantiates a new GetUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsage200Response() *GetUsage200Response {
	this := GetUsage200Response{}
	return &this
}

// NewGetUsage200ResponseWithDefaults instantiates a new GetUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsage200ResponseWithDefaults() *GetUsage200Response {
	this := GetUsage200Response{}
	return &this
}

// GetCharacterCount returns the CharacterCount field value if set, zero value otherwise.
func (o *GetUsage200Response) GetCharacterCount() int64 {
	if o == nil || o.CharacterCount == nil {
		var ret int64
		return ret
	}
	return *o.CharacterCount
}

// GetCharacterCountOk returns a tuple with the CharacterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsage200Response) GetCharacterCountOk() (*int64, bool) {
	if o == nil || o.CharacterCount == nil {
		return nil, false
	}
	return o.CharacterCount, true
}

// HasCharacterCount returns a boolean if a field has been set.
func (o *GetUsage200Response) HasCharacterCount() bool {
	if o != nil && o.CharacterCount != nil {
		return true
	}

	return false
}

// SetCharacterCount gets a reference to the given int64 and assigns it to the CharacterCount field.
func (o *GetUsage200Response) SetCharacterCount(v int64) {
	o.CharacterCount = &v
}

// GetCharacterLimit returns the CharacterLimit field value if set, zero value otherwise.
func (o *GetUsage200Response) GetCharacterLimit() int64 {
	if o == nil || o.CharacterLimit == nil {
		var ret int64
		return ret
	}
	return *o.CharacterLimit
}

// GetCharacterLimitOk returns a tuple with the CharacterLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsage200Response) GetCharacterLimitOk() (*int64, bool) {
	if o == nil || o.CharacterLimit == nil {
		return nil, false
	}
	return o.CharacterLimit, true
}

// HasCharacterLimit returns a boolean if a field has been set.
func (o *GetUsage200Response) HasCharacterLimit() bool {
	if o != nil && o.CharacterLimit != nil {
		return true
	}

	return false
}

// SetCharacterLimit gets a reference to the given int64 and assigns it to the CharacterLimit field.
func (o *GetUsage200Response) SetCharacterLimit(v int64) {
	o.CharacterLimit = &v
}

// GetDocumentLimit returns the DocumentLimit field value if set, zero value otherwise.
func (o *GetUsage200Response) GetDocumentLimit() int64 {
	if o == nil || o.DocumentLimit == nil {
		var ret int64
		return ret
	}
	return *o.DocumentLimit
}

// GetDocumentLimitOk returns a tuple with the DocumentLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsage200Response) GetDocumentLimitOk() (*int64, bool) {
	if o == nil || o.DocumentLimit == nil {
		return nil, false
	}
	return o.DocumentLimit, true
}

// HasDocumentLimit returns a boolean if a field has been set.
func (o *GetUsage200Response) HasDocumentLimit() bool {
	if o != nil && o.DocumentLimit != nil {
		return true
	}

	return false
}

// SetDocumentLimit gets a reference to the given int64 and assigns it to the DocumentLimit field.
func (o *GetUsage200Response) SetDocumentLimit(v int64) {
	o.DocumentLimit = &v
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise.
func (o *GetUsage200Response) GetDocumentCount() int64 {
	if o == nil || o.DocumentCount == nil {
		var ret int64
		return ret
	}
	return *o.DocumentCount
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsage200Response) GetDocumentCountOk() (*int64, bool) {
	if o == nil || o.DocumentCount == nil {
		return nil, false
	}
	return o.DocumentCount, true
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *GetUsage200Response) HasDocumentCount() bool {
	if o != nil && o.DocumentCount != nil {
		return true
	}

	return false
}

// SetDocumentCount gets a reference to the given int64 and assigns it to the DocumentCount field.
func (o *GetUsage200Response) SetDocumentCount(v int64) {
	o.DocumentCount = &v
}

// GetTeamDocumentLimit returns the TeamDocumentLimit field value if set, zero value otherwise.
func (o *GetUsage200Response) GetTeamDocumentLimit() int64 {
	if o == nil || o.TeamDocumentLimit == nil {
		var ret int64
		return ret
	}
	return *o.TeamDocumentLimit
}

// GetTeamDocumentLimitOk returns a tuple with the TeamDocumentLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsage200Response) GetTeamDocumentLimitOk() (*int64, bool) {
	if o == nil || o.TeamDocumentLimit == nil {
		return nil, false
	}
	return o.TeamDocumentLimit, true
}

// HasTeamDocumentLimit returns a boolean if a field has been set.
func (o *GetUsage200Response) HasTeamDocumentLimit() bool {
	if o != nil && o.TeamDocumentLimit != nil {
		return true
	}

	return false
}

// SetTeamDocumentLimit gets a reference to the given int64 and assigns it to the TeamDocumentLimit field.
func (o *GetUsage200Response) SetTeamDocumentLimit(v int64) {
	o.TeamDocumentLimit = &v
}

// GetTeamDocumentCount returns the TeamDocumentCount field value if set, zero value otherwise.
func (o *GetUsage200Response) GetTeamDocumentCount() int64 {
	if o == nil || o.TeamDocumentCount == nil {
		var ret int64
		return ret
	}
	return *o.TeamDocumentCount
}

// GetTeamDocumentCountOk returns a tuple with the TeamDocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsage200Response) GetTeamDocumentCountOk() (*int64, bool) {
	if o == nil || o.TeamDocumentCount == nil {
		return nil, false
	}
	return o.TeamDocumentCount, true
}

// HasTeamDocumentCount returns a boolean if a field has been set.
func (o *GetUsage200Response) HasTeamDocumentCount() bool {
	if o != nil && o.TeamDocumentCount != nil {
		return true
	}

	return false
}

// SetTeamDocumentCount gets a reference to the given int64 and assigns it to the TeamDocumentCount field.
func (o *GetUsage200Response) SetTeamDocumentCount(v int64) {
	o.TeamDocumentCount = &v
}

func (o GetUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CharacterCount != nil {
		toSerialize["character_count"] = o.CharacterCount
	}
	if o.CharacterLimit != nil {
		toSerialize["character_limit"] = o.CharacterLimit
	}
	if o.DocumentLimit != nil {
		toSerialize["document_limit"] = o.DocumentLimit
	}
	if o.DocumentCount != nil {
		toSerialize["document_count"] = o.DocumentCount
	}
	if o.TeamDocumentLimit != nil {
		toSerialize["team_document_limit"] = o.TeamDocumentLimit
	}
	if o.TeamDocumentCount != nil {
		toSerialize["team_document_count"] = o.TeamDocumentCount
	}
	return json.Marshal(toSerialize)
}

type NullableGetUsage200Response struct {
	value *GetUsage200Response
	isSet bool
}

func (v NullableGetUsage200Response) Get() *GetUsage200Response {
	return v.value
}

func (v *NullableGetUsage200Response) Set(val *GetUsage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsage200Response(val *GetUsage200Response) *NullableGetUsage200Response {
	return &NullableGetUsage200Response{value: val, isSet: true}
}

func (v NullableGetUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

