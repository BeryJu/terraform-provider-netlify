/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
)

// checks if the MemberCommitterMatchMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberCommitterMatchMethod{}

// MemberCommitterMatchMethod the method the member is matched to a committer by.
type MemberCommitterMatchMethod struct {
	Automatic *string `json:"automatic,omitempty"`
	Manual *string `json:"manual,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberCommitterMatchMethod MemberCommitterMatchMethod

// NewMemberCommitterMatchMethod instantiates a new MemberCommitterMatchMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberCommitterMatchMethod() *MemberCommitterMatchMethod {
	this := MemberCommitterMatchMethod{}
	return &this
}

// NewMemberCommitterMatchMethodWithDefaults instantiates a new MemberCommitterMatchMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberCommitterMatchMethodWithDefaults() *MemberCommitterMatchMethod {
	this := MemberCommitterMatchMethod{}
	return &this
}

// GetAutomatic returns the Automatic field value if set, zero value otherwise.
func (o *MemberCommitterMatchMethod) GetAutomatic() string {
	if o == nil || IsNil(o.Automatic) {
		var ret string
		return ret
	}
	return *o.Automatic
}

// GetAutomaticOk returns a tuple with the Automatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCommitterMatchMethod) GetAutomaticOk() (*string, bool) {
	if o == nil || IsNil(o.Automatic) {
		return nil, false
	}
	return o.Automatic, true
}

// HasAutomatic returns a boolean if a field has been set.
func (o *MemberCommitterMatchMethod) HasAutomatic() bool {
	if o != nil && !IsNil(o.Automatic) {
		return true
	}

	return false
}

// SetAutomatic gets a reference to the given string and assigns it to the Automatic field.
func (o *MemberCommitterMatchMethod) SetAutomatic(v string) {
	o.Automatic = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *MemberCommitterMatchMethod) GetManual() string {
	if o == nil || IsNil(o.Manual) {
		var ret string
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCommitterMatchMethod) GetManualOk() (*string, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *MemberCommitterMatchMethod) HasManual() bool {
	if o != nil && !IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given string and assigns it to the Manual field.
func (o *MemberCommitterMatchMethod) SetManual(v string) {
	o.Manual = &v
}

func (o MemberCommitterMatchMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberCommitterMatchMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Automatic) {
		toSerialize["automatic"] = o.Automatic
	}
	if !IsNil(o.Manual) {
		toSerialize["manual"] = o.Manual
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberCommitterMatchMethod) UnmarshalJSON(data []byte) (err error) {
	varMemberCommitterMatchMethod := _MemberCommitterMatchMethod{}

	err = json.Unmarshal(data, &varMemberCommitterMatchMethod)

	if err != nil {
		return err
	}

	*o = MemberCommitterMatchMethod(varMemberCommitterMatchMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "automatic")
		delete(additionalProperties, "manual")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberCommitterMatchMethod struct {
	value *MemberCommitterMatchMethod
	isSet bool
}

func (v NullableMemberCommitterMatchMethod) Get() *MemberCommitterMatchMethod {
	return v.value
}

func (v *NullableMemberCommitterMatchMethod) Set(val *MemberCommitterMatchMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberCommitterMatchMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberCommitterMatchMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberCommitterMatchMethod(val *MemberCommitterMatchMethod) *NullableMemberCommitterMatchMethod {
	return &NullableMemberCommitterMatchMethod{value: val, isSet: true}
}

func (v NullableMemberCommitterMatchMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberCommitterMatchMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

