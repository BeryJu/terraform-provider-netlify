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

// checks if the EnvVarUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvVarUser{}

// EnvVarUser struct for EnvVarUser
type EnvVarUser struct {
	// The user's unique identifier
	Id *string `json:"id,omitempty"`
	// The user's full name (first and last)
	FullName *string `json:"full_name,omitempty"`
	// The user's email address
	Email *string `json:"email,omitempty"`
	// A URL pointing to the user's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvVarUser EnvVarUser

// NewEnvVarUser instantiates a new EnvVarUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvVarUser() *EnvVarUser {
	this := EnvVarUser{}
	return &this
}

// NewEnvVarUserWithDefaults instantiates a new EnvVarUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvVarUserWithDefaults() *EnvVarUser {
	this := EnvVarUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvVarUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvVarUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvVarUser) SetId(v string) {
	o.Id = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *EnvVarUser) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarUser) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *EnvVarUser) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *EnvVarUser) SetFullName(v string) {
	o.FullName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EnvVarUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EnvVarUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EnvVarUser) SetEmail(v string) {
	o.Email = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *EnvVarUser) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *EnvVarUser) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *EnvVarUser) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

func (o EnvVarUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvVarUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvVarUser) UnmarshalJSON(data []byte) (err error) {
	varEnvVarUser := _EnvVarUser{}

	err = json.Unmarshal(data, &varEnvVarUser)

	if err != nil {
		return err
	}

	*o = EnvVarUser(varEnvVarUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "avatar_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvVarUser struct {
	value *EnvVarUser
	isSet bool
}

func (v NullableEnvVarUser) Get() *EnvVarUser {
	return v.value
}

func (v *NullableEnvVarUser) Set(val *EnvVarUser) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvVarUser) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvVarUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvVarUser(val *EnvVarUser) *NullableEnvVarUser {
	return &NullableEnvVarUser{value: val, isSet: true}
}

func (v NullableEnvVarUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvVarUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

