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

// checks if the IdentityVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityVerification{}

// IdentityVerification struct for IdentityVerification
type IdentityVerification struct {
	// The provider of the IdentityVerification (stripe)
	Provider *string `json:"provider,omitempty"`
	// The UUID for IdentityVerification lookups
	Code *string `json:"code,omitempty"`
	// The session_id for IdentityVerification session
	SessionId *string `json:"session_id,omitempty"`
	// The report from the identity verification service
	Report map[string]interface{} `json:"report,omitempty"`
	User *User `json:"user,omitempty"`
	// The current state of the verification
	State *string `json:"state,omitempty"`
	// When the Event IdentityVerification was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// When the Event IdentityVerification was updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerification IdentityVerification

// NewIdentityVerification instantiates a new IdentityVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerification() *IdentityVerification {
	this := IdentityVerification{}
	return &this
}

// NewIdentityVerificationWithDefaults instantiates a new IdentityVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationWithDefaults() *IdentityVerification {
	this := IdentityVerification{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *IdentityVerification) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *IdentityVerification) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *IdentityVerification) SetProvider(v string) {
	o.Provider = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IdentityVerification) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IdentityVerification) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IdentityVerification) SetCode(v string) {
	o.Code = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *IdentityVerification) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *IdentityVerification) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *IdentityVerification) SetSessionId(v string) {
	o.SessionId = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *IdentityVerification) GetReport() map[string]interface{} {
	if o == nil || IsNil(o.Report) {
		var ret map[string]interface{}
		return ret
	}
	return o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetReportOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Report) {
		return map[string]interface{}{}, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *IdentityVerification) HasReport() bool {
	if o != nil && !IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given map[string]interface{} and assigns it to the Report field.
func (o *IdentityVerification) SetReport(v map[string]interface{}) {
	o.Report = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IdentityVerification) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IdentityVerification) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *IdentityVerification) SetUser(v User) {
	o.User = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IdentityVerification) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IdentityVerification) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IdentityVerification) SetState(v string) {
	o.State = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityVerification) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityVerification) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IdentityVerification) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityVerification) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityVerification) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IdentityVerification) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IdentityVerification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.SessionId) {
		toSerialize["session_id"] = o.SessionId
	}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityVerification) UnmarshalJSON(data []byte) (err error) {
	varIdentityVerification := _IdentityVerification{}

	err = json.Unmarshal(data, &varIdentityVerification)

	if err != nil {
		return err
	}

	*o = IdentityVerification(varIdentityVerification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider")
		delete(additionalProperties, "code")
		delete(additionalProperties, "session_id")
		delete(additionalProperties, "report")
		delete(additionalProperties, "user")
		delete(additionalProperties, "state")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerification struct {
	value *IdentityVerification
	isSet bool
}

func (v NullableIdentityVerification) Get() *IdentityVerification {
	return v.value
}

func (v *NullableIdentityVerification) Set(val *IdentityVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerification(val *IdentityVerification) *NullableIdentityVerification {
	return &NullableIdentityVerification{value: val, isSet: true}
}

func (v NullableIdentityVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

