/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
	"time"
)

// checks if the AccountBuildPriority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountBuildPriority{}

// AccountBuildPriority included if the build is prioritized
type AccountBuildPriority struct {
	PrioritizedAt *time.Time `json:"prioritized_at,omitempty"`
	PrioritizedBy *string `json:"prioritized_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBuildPriority AccountBuildPriority

// NewAccountBuildPriority instantiates a new AccountBuildPriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBuildPriority() *AccountBuildPriority {
	this := AccountBuildPriority{}
	return &this
}

// NewAccountBuildPriorityWithDefaults instantiates a new AccountBuildPriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBuildPriorityWithDefaults() *AccountBuildPriority {
	this := AccountBuildPriority{}
	return &this
}

// GetPrioritizedAt returns the PrioritizedAt field value if set, zero value otherwise.
func (o *AccountBuildPriority) GetPrioritizedAt() time.Time {
	if o == nil || IsNil(o.PrioritizedAt) {
		var ret time.Time
		return ret
	}
	return *o.PrioritizedAt
}

// GetPrioritizedAtOk returns a tuple with the PrioritizedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBuildPriority) GetPrioritizedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PrioritizedAt) {
		return nil, false
	}
	return o.PrioritizedAt, true
}

// HasPrioritizedAt returns a boolean if a field has been set.
func (o *AccountBuildPriority) HasPrioritizedAt() bool {
	if o != nil && !IsNil(o.PrioritizedAt) {
		return true
	}

	return false
}

// SetPrioritizedAt gets a reference to the given time.Time and assigns it to the PrioritizedAt field.
func (o *AccountBuildPriority) SetPrioritizedAt(v time.Time) {
	o.PrioritizedAt = &v
}

// GetPrioritizedBy returns the PrioritizedBy field value if set, zero value otherwise.
func (o *AccountBuildPriority) GetPrioritizedBy() string {
	if o == nil || IsNil(o.PrioritizedBy) {
		var ret string
		return ret
	}
	return *o.PrioritizedBy
}

// GetPrioritizedByOk returns a tuple with the PrioritizedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBuildPriority) GetPrioritizedByOk() (*string, bool) {
	if o == nil || IsNil(o.PrioritizedBy) {
		return nil, false
	}
	return o.PrioritizedBy, true
}

// HasPrioritizedBy returns a boolean if a field has been set.
func (o *AccountBuildPriority) HasPrioritizedBy() bool {
	if o != nil && !IsNil(o.PrioritizedBy) {
		return true
	}

	return false
}

// SetPrioritizedBy gets a reference to the given string and assigns it to the PrioritizedBy field.
func (o *AccountBuildPriority) SetPrioritizedBy(v string) {
	o.PrioritizedBy = &v
}

func (o AccountBuildPriority) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBuildPriority) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrioritizedAt) {
		toSerialize["prioritized_at"] = o.PrioritizedAt
	}
	if !IsNil(o.PrioritizedBy) {
		toSerialize["prioritized_by"] = o.PrioritizedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountBuildPriority) UnmarshalJSON(data []byte) (err error) {
	varAccountBuildPriority := _AccountBuildPriority{}

	err = json.Unmarshal(data, &varAccountBuildPriority)

	if err != nil {
		return err
	}

	*o = AccountBuildPriority(varAccountBuildPriority)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prioritized_at")
		delete(additionalProperties, "prioritized_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBuildPriority struct {
	value *AccountBuildPriority
	isSet bool
}

func (v NullableAccountBuildPriority) Get() *AccountBuildPriority {
	return v.value
}

func (v *NullableAccountBuildPriority) Set(val *AccountBuildPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBuildPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBuildPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBuildPriority(val *AccountBuildPriority) *NullableAccountBuildPriority {
	return &NullableAccountBuildPriority{value: val, isSet: true}
}

func (v NullableAccountBuildPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBuildPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

