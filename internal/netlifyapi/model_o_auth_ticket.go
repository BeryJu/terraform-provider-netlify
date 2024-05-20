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
	"fmt"
)

// checks if the OAuthTicket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthTicket{}

// OAuthTicket struct for OAuthTicket
type OAuthTicket struct {
	// The ID for the OAuth ticket
	Id string `json:"id"`
	// The UID of the OAuth application
	ClientId string `json:"client_id"`
	// Whether the OAuth ticket is authorized
	Authorized bool `json:"authorized"`
	// When the OAuth ticket was created
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _OAuthTicket OAuthTicket

// NewOAuthTicket instantiates a new OAuthTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthTicket(id string, clientId string, authorized bool, createdAt time.Time) *OAuthTicket {
	this := OAuthTicket{}
	this.Id = id
	this.ClientId = clientId
	this.Authorized = authorized
	this.CreatedAt = createdAt
	return &this
}

// NewOAuthTicketWithDefaults instantiates a new OAuthTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthTicketWithDefaults() *OAuthTicket {
	this := OAuthTicket{}
	return &this
}

// GetId returns the Id field value
func (o *OAuthTicket) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OAuthTicket) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OAuthTicket) SetId(v string) {
	o.Id = v
}

// GetClientId returns the ClientId field value
func (o *OAuthTicket) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuthTicket) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuthTicket) SetClientId(v string) {
	o.ClientId = v
}

// GetAuthorized returns the Authorized field value
func (o *OAuthTicket) GetAuthorized() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Authorized
}

// GetAuthorizedOk returns a tuple with the Authorized field value
// and a boolean to check if the value has been set.
func (o *OAuthTicket) GetAuthorizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorized, true
}

// SetAuthorized sets field value
func (o *OAuthTicket) SetAuthorized(v bool) {
	o.Authorized = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OAuthTicket) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OAuthTicket) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OAuthTicket) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o OAuthTicket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthTicket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["client_id"] = o.ClientId
	toSerialize["authorized"] = o.Authorized
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuthTicket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"client_id",
		"authorized",
		"created_at",
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

	varOAuthTicket := _OAuthTicket{}

	err = json.Unmarshal(data, &varOAuthTicket)

	if err != nil {
		return err
	}

	*o = OAuthTicket(varOAuthTicket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "authorized")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuthTicket struct {
	value *OAuthTicket
	isSet bool
}

func (v NullableOAuthTicket) Get() *OAuthTicket {
	return v.value
}

func (v *NullableOAuthTicket) Set(val *OAuthTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthTicket(val *OAuthTicket) *NullableOAuthTicket {
	return &NullableOAuthTicket{value: val, isSet: true}
}

func (v NullableOAuthTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


