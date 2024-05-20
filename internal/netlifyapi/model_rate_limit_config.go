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

// checks if the RateLimitConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitConfig{}

// RateLimitConfig struct for RateLimitConfig
type RateLimitConfig struct {
	// ]
	Algorithm string `json:"algorithm"`
	// In seconds
	WindowSize *int64 `json:"window_size,omitempty"`
	WindowLimit int64 `json:"window_limit"`
	AdditionalProperties map[string]interface{}
}

type _RateLimitConfig RateLimitConfig

// NewRateLimitConfig instantiates a new RateLimitConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitConfig(algorithm string, windowLimit int64) *RateLimitConfig {
	this := RateLimitConfig{}
	this.Algorithm = algorithm
	this.WindowLimit = windowLimit
	return &this
}

// NewRateLimitConfigWithDefaults instantiates a new RateLimitConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitConfigWithDefaults() *RateLimitConfig {
	this := RateLimitConfig{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *RateLimitConfig) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *RateLimitConfig) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *RateLimitConfig) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *RateLimitConfig) GetWindowSize() int64 {
	if o == nil || IsNil(o.WindowSize) {
		var ret int64
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitConfig) GetWindowSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.WindowSize) {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *RateLimitConfig) HasWindowSize() bool {
	if o != nil && !IsNil(o.WindowSize) {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int64 and assigns it to the WindowSize field.
func (o *RateLimitConfig) SetWindowSize(v int64) {
	o.WindowSize = &v
}

// GetWindowLimit returns the WindowLimit field value
func (o *RateLimitConfig) GetWindowLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WindowLimit
}

// GetWindowLimitOk returns a tuple with the WindowLimit field value
// and a boolean to check if the value has been set.
func (o *RateLimitConfig) GetWindowLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowLimit, true
}

// SetWindowLimit sets field value
func (o *RateLimitConfig) SetWindowLimit(v int64) {
	o.WindowLimit = v
}

func (o RateLimitConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	if !IsNil(o.WindowSize) {
		toSerialize["window_size"] = o.WindowSize
	}
	toSerialize["window_limit"] = o.WindowLimit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RateLimitConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algorithm",
		"window_limit",
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

	varRateLimitConfig := _RateLimitConfig{}

	err = json.Unmarshal(data, &varRateLimitConfig)

	if err != nil {
		return err
	}

	*o = RateLimitConfig(varRateLimitConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "window_size")
		delete(additionalProperties, "window_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRateLimitConfig struct {
	value *RateLimitConfig
	isSet bool
}

func (v NullableRateLimitConfig) Get() *RateLimitConfig {
	return v.value
}

func (v *NullableRateLimitConfig) Set(val *RateLimitConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitConfig(val *RateLimitConfig) *NullableRateLimitConfig {
	return &NullableRateLimitConfig{value: val, isSet: true}
}

func (v NullableRateLimitConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


