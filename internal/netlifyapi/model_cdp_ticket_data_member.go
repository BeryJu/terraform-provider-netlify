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

// checks if the CDPTicketDataMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CDPTicketDataMember{}

// CDPTicketDataMember struct for CDPTicketDataMember
type CDPTicketDataMember struct {
	Avatar string `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	Disabled bool `json:"disabled"`
	Email string `json:"email"`
	Id string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Username string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _CDPTicketDataMember CDPTicketDataMember

// NewCDPTicketDataMember instantiates a new CDPTicketDataMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCDPTicketDataMember(avatar string, createdAt time.Time, disabled bool, email string, id string, name string, role string, username string) *CDPTicketDataMember {
	this := CDPTicketDataMember{}
	this.Avatar = avatar
	this.CreatedAt = createdAt
	this.Disabled = disabled
	this.Email = email
	this.Id = id
	this.Name = name
	this.Role = role
	this.Username = username
	return &this
}

// NewCDPTicketDataMemberWithDefaults instantiates a new CDPTicketDataMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCDPTicketDataMemberWithDefaults() *CDPTicketDataMember {
	this := CDPTicketDataMember{}
	return &this
}

// GetAvatar returns the Avatar field value
func (o *CDPTicketDataMember) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *CDPTicketDataMember) SetAvatar(v string) {
	o.Avatar = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CDPTicketDataMember) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CDPTicketDataMember) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDisabled returns the Disabled field value
func (o *CDPTicketDataMember) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *CDPTicketDataMember) SetDisabled(v bool) {
	o.Disabled = v
}

// GetEmail returns the Email field value
func (o *CDPTicketDataMember) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CDPTicketDataMember) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *CDPTicketDataMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CDPTicketDataMember) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CDPTicketDataMember) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CDPTicketDataMember) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *CDPTicketDataMember) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CDPTicketDataMember) SetRole(v string) {
	o.Role = v
}

// GetUsername returns the Username field value
func (o *CDPTicketDataMember) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataMember) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CDPTicketDataMember) SetUsername(v string) {
	o.Username = v
}

func (o CDPTicketDataMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CDPTicketDataMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["avatar"] = o.Avatar
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["disabled"] = o.Disabled
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CDPTicketDataMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"avatar",
		"created_at",
		"disabled",
		"email",
		"id",
		"name",
		"role",
		"username",
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

	varCDPTicketDataMember := _CDPTicketDataMember{}

	err = json.Unmarshal(data, &varCDPTicketDataMember)

	if err != nil {
		return err
	}

	*o = CDPTicketDataMember(varCDPTicketDataMember)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avatar")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCDPTicketDataMember struct {
	value *CDPTicketDataMember
	isSet bool
}

func (v NullableCDPTicketDataMember) Get() *CDPTicketDataMember {
	return v.value
}

func (v *NullableCDPTicketDataMember) Set(val *CDPTicketDataMember) {
	v.value = val
	v.isSet = true
}

func (v NullableCDPTicketDataMember) IsSet() bool {
	return v.isSet
}

func (v *NullableCDPTicketDataMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCDPTicketDataMember(val *CDPTicketDataMember) *NullableCDPTicketDataMember {
	return &NullableCDPTicketDataMember{value: val, isSet: true}
}

func (v NullableCDPTicketDataMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCDPTicketDataMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


