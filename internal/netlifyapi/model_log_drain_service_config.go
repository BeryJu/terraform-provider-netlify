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

// checks if the LogDrainServiceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogDrainServiceConfig{}

// LogDrainServiceConfig The configuration for sending log drains to the destination service
type LogDrainServiceConfig struct {
	// The URL for the log drain
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _LogDrainServiceConfig LogDrainServiceConfig

// NewLogDrainServiceConfig instantiates a new LogDrainServiceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDrainServiceConfig(url string) *LogDrainServiceConfig {
	this := LogDrainServiceConfig{}
	this.Url = url
	return &this
}

// NewLogDrainServiceConfigWithDefaults instantiates a new LogDrainServiceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDrainServiceConfigWithDefaults() *LogDrainServiceConfig {
	this := LogDrainServiceConfig{}
	return &this
}

// GetUrl returns the Url field value
func (o *LogDrainServiceConfig) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LogDrainServiceConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LogDrainServiceConfig) SetUrl(v string) {
	o.Url = v
}

func (o LogDrainServiceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogDrainServiceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogDrainServiceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varLogDrainServiceConfig := _LogDrainServiceConfig{}

	err = json.Unmarshal(data, &varLogDrainServiceConfig)

	if err != nil {
		return err
	}

	*o = LogDrainServiceConfig(varLogDrainServiceConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogDrainServiceConfig struct {
	value *LogDrainServiceConfig
	isSet bool
}

func (v NullableLogDrainServiceConfig) Get() *LogDrainServiceConfig {
	return v.value
}

func (v *NullableLogDrainServiceConfig) Set(val *LogDrainServiceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDrainServiceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDrainServiceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDrainServiceConfig(val *LogDrainServiceConfig) *NullableLogDrainServiceConfig {
	return &NullableLogDrainServiceConfig{value: val, isSet: true}
}

func (v NullableLogDrainServiceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDrainServiceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


