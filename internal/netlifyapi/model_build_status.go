/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
	"fmt"
)

// checks if the BuildStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildStatus{}

// BuildStatus struct for BuildStatus
type BuildStatus struct {
	// Number of active builds
	Active int64 `json:"active"`
	// Number of pending concurrency
	PendingConcurrency int64 `json:"pending_concurrency"`
	// Number of enqueued builds
	Enqueued int64 `json:"enqueued"`
	Minutes BuildStatusMinutes `json:"minutes"`
	AdditionalProperties map[string]interface{}
}

type _BuildStatus BuildStatus

// NewBuildStatus instantiates a new BuildStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildStatus(active int64, pendingConcurrency int64, enqueued int64, minutes BuildStatusMinutes) *BuildStatus {
	this := BuildStatus{}
	this.Active = active
	this.PendingConcurrency = pendingConcurrency
	this.Enqueued = enqueued
	this.Minutes = minutes
	return &this
}

// NewBuildStatusWithDefaults instantiates a new BuildStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildStatusWithDefaults() *BuildStatus {
	this := BuildStatus{}
	return &this
}

// GetActive returns the Active field value
func (o *BuildStatus) GetActive() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *BuildStatus) GetActiveOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *BuildStatus) SetActive(v int64) {
	o.Active = v
}

// GetPendingConcurrency returns the PendingConcurrency field value
func (o *BuildStatus) GetPendingConcurrency() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PendingConcurrency
}

// GetPendingConcurrencyOk returns a tuple with the PendingConcurrency field value
// and a boolean to check if the value has been set.
func (o *BuildStatus) GetPendingConcurrencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingConcurrency, true
}

// SetPendingConcurrency sets field value
func (o *BuildStatus) SetPendingConcurrency(v int64) {
	o.PendingConcurrency = v
}

// GetEnqueued returns the Enqueued field value
func (o *BuildStatus) GetEnqueued() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Enqueued
}

// GetEnqueuedOk returns a tuple with the Enqueued field value
// and a boolean to check if the value has been set.
func (o *BuildStatus) GetEnqueuedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enqueued, true
}

// SetEnqueued sets field value
func (o *BuildStatus) SetEnqueued(v int64) {
	o.Enqueued = v
}

// GetMinutes returns the Minutes field value
func (o *BuildStatus) GetMinutes() BuildStatusMinutes {
	if o == nil {
		var ret BuildStatusMinutes
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *BuildStatus) GetMinutesOk() (*BuildStatusMinutes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *BuildStatus) SetMinutes(v BuildStatusMinutes) {
	o.Minutes = v
}

func (o BuildStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["pending_concurrency"] = o.PendingConcurrency
	toSerialize["enqueued"] = o.Enqueued
	toSerialize["minutes"] = o.Minutes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BuildStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"pending_concurrency",
		"enqueued",
		"minutes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBuildStatus := _BuildStatus{}

	err = json.Unmarshal(data, &varBuildStatus)

	if err != nil {
		return err
	}

	*o = BuildStatus(varBuildStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "pending_concurrency")
		delete(additionalProperties, "enqueued")
		delete(additionalProperties, "minutes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildStatus struct {
	value *BuildStatus
	isSet bool
}

func (v NullableBuildStatus) Get() *BuildStatus {
	return v.value
}

func (v *NullableBuildStatus) Set(val *BuildStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildStatus(val *BuildStatus) *NullableBuildStatus {
	return &NullableBuildStatus{value: val, isSet: true}
}

func (v NullableBuildStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

