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

// checks if the GeoTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoTargetConfig{}

// GeoTargetConfig struct for GeoTargetConfig
type GeoTargetConfig struct {
	Country string `json:"country"`
	SubRegion *string `json:"sub_region,omitempty"`
	Exclude *bool `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GeoTargetConfig GeoTargetConfig

// NewGeoTargetConfig instantiates a new GeoTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoTargetConfig(country string) *GeoTargetConfig {
	this := GeoTargetConfig{}
	this.Country = country
	return &this
}

// NewGeoTargetConfigWithDefaults instantiates a new GeoTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoTargetConfigWithDefaults() *GeoTargetConfig {
	this := GeoTargetConfig{}
	return &this
}

// GetCountry returns the Country field value
func (o *GeoTargetConfig) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *GeoTargetConfig) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *GeoTargetConfig) SetCountry(v string) {
	o.Country = v
}

// GetSubRegion returns the SubRegion field value if set, zero value otherwise.
func (o *GeoTargetConfig) GetSubRegion() string {
	if o == nil || IsNil(o.SubRegion) {
		var ret string
		return ret
	}
	return *o.SubRegion
}

// GetSubRegionOk returns a tuple with the SubRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoTargetConfig) GetSubRegionOk() (*string, bool) {
	if o == nil || IsNil(o.SubRegion) {
		return nil, false
	}
	return o.SubRegion, true
}

// HasSubRegion returns a boolean if a field has been set.
func (o *GeoTargetConfig) HasSubRegion() bool {
	if o != nil && !IsNil(o.SubRegion) {
		return true
	}

	return false
}

// SetSubRegion gets a reference to the given string and assigns it to the SubRegion field.
func (o *GeoTargetConfig) SetSubRegion(v string) {
	o.SubRegion = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *GeoTargetConfig) GetExclude() bool {
	if o == nil || IsNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoTargetConfig) GetExcludeOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *GeoTargetConfig) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *GeoTargetConfig) SetExclude(v bool) {
	o.Exclude = &v
}

func (o GeoTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	if !IsNil(o.SubRegion) {
		toSerialize["sub_region"] = o.SubRegion
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GeoTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country",
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

	varGeoTargetConfig := _GeoTargetConfig{}

	err = json.Unmarshal(data, &varGeoTargetConfig)

	if err != nil {
		return err
	}

	*o = GeoTargetConfig(varGeoTargetConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "country")
		delete(additionalProperties, "sub_region")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGeoTargetConfig struct {
	value *GeoTargetConfig
	isSet bool
}

func (v NullableGeoTargetConfig) Get() *GeoTargetConfig {
	return v.value
}

func (v *NullableGeoTargetConfig) Set(val *GeoTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoTargetConfig(val *GeoTargetConfig) *NullableGeoTargetConfig {
	return &NullableGeoTargetConfig{value: val, isSet: true}
}

func (v NullableGeoTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

