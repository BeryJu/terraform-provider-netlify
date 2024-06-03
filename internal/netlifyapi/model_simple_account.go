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

// checks if the SimpleAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleAccount{}

// SimpleAccount SimpleAccount model definition
type SimpleAccount struct {
	// The ID of the account
	Id string `json:"id"`
	// The name of the account
	Name string `json:"name"`
	// The URL of the team logo
	TeamLogoUrl string `json:"team_logo_url"`
	// The hash list of capabilities related to CDP
	CdpCapabilities map[string]interface{} `json:"cdp_capabilities"`
	AdditionalProperties map[string]interface{}
}

type _SimpleAccount SimpleAccount

// NewSimpleAccount instantiates a new SimpleAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleAccount(id string, name string, teamLogoUrl string, cdpCapabilities map[string]interface{}) *SimpleAccount {
	this := SimpleAccount{}
	this.Id = id
	this.Name = name
	this.TeamLogoUrl = teamLogoUrl
	this.CdpCapabilities = cdpCapabilities
	return &this
}

// NewSimpleAccountWithDefaults instantiates a new SimpleAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleAccountWithDefaults() *SimpleAccount {
	this := SimpleAccount{}
	return &this
}

// GetId returns the Id field value
func (o *SimpleAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SimpleAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SimpleAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SimpleAccount) SetName(v string) {
	o.Name = v
}

// GetTeamLogoUrl returns the TeamLogoUrl field value
func (o *SimpleAccount) GetTeamLogoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamLogoUrl
}

// GetTeamLogoUrlOk returns a tuple with the TeamLogoUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleAccount) GetTeamLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamLogoUrl, true
}

// SetTeamLogoUrl sets field value
func (o *SimpleAccount) SetTeamLogoUrl(v string) {
	o.TeamLogoUrl = v
}

// GetCdpCapabilities returns the CdpCapabilities field value
func (o *SimpleAccount) GetCdpCapabilities() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CdpCapabilities
}

// GetCdpCapabilitiesOk returns a tuple with the CdpCapabilities field value
// and a boolean to check if the value has been set.
func (o *SimpleAccount) GetCdpCapabilitiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CdpCapabilities, true
}

// SetCdpCapabilities sets field value
func (o *SimpleAccount) SetCdpCapabilities(v map[string]interface{}) {
	o.CdpCapabilities = v
}

func (o SimpleAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["team_logo_url"] = o.TeamLogoUrl
	toSerialize["cdp_capabilities"] = o.CdpCapabilities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimpleAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"team_logo_url",
		"cdp_capabilities",
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

	varSimpleAccount := _SimpleAccount{}

	err = json.Unmarshal(data, &varSimpleAccount)

	if err != nil {
		return err
	}

	*o = SimpleAccount(varSimpleAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "team_logo_url")
		delete(additionalProperties, "cdp_capabilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimpleAccount struct {
	value *SimpleAccount
	isSet bool
}

func (v NullableSimpleAccount) Get() *SimpleAccount {
	return v.value
}

func (v *NullableSimpleAccount) Set(val *SimpleAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleAccount(val *SimpleAccount) *NullableSimpleAccount {
	return &NullableSimpleAccount{value: val, isSet: true}
}

func (v NullableSimpleAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

