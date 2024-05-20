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

// checks if the Reviewer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reviewer{}

// Reviewer Reviewer model definition
type Reviewer struct {
	// The ID for the reviewer
	Id string `json:"id"`
	// The ID of the account the reviewer is associated with
	AccountId string `json:"account_id"`
	// The ID of the user associated with the reviewer
	UserId string `json:"user_id"`
	// The state of the reviewer
	State string `json:"state"`
	// The site access type of the reviewer
	SiteAccess string `json:"site_access"`
	// The email address of the reviewer
	Email string `json:"email"`
	// The full name of the reviewer
	FullName string `json:"full_name"`
	// When the reviewer was created
	CreatedAt time.Time `json:"created_at"`
	// Last time the reviewer was updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Reviewer Reviewer

// NewReviewer instantiates a new Reviewer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewer(id string, accountId string, userId string, state string, siteAccess string, email string, fullName string, createdAt time.Time, updatedAt time.Time) *Reviewer {
	this := Reviewer{}
	this.Id = id
	this.AccountId = accountId
	this.UserId = userId
	this.State = state
	this.SiteAccess = siteAccess
	this.Email = email
	this.FullName = fullName
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewReviewerWithDefaults instantiates a new Reviewer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerWithDefaults() *Reviewer {
	this := Reviewer{}
	return &this
}

// GetId returns the Id field value
func (o *Reviewer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Reviewer) SetId(v string) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *Reviewer) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Reviewer) SetAccountId(v string) {
	o.AccountId = v
}

// GetUserId returns the UserId field value
func (o *Reviewer) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Reviewer) SetUserId(v string) {
	o.UserId = v
}

// GetState returns the State field value
func (o *Reviewer) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Reviewer) SetState(v string) {
	o.State = v
}

// GetSiteAccess returns the SiteAccess field value
func (o *Reviewer) GetSiteAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteAccess
}

// GetSiteAccessOk returns a tuple with the SiteAccess field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetSiteAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteAccess, true
}

// SetSiteAccess sets field value
func (o *Reviewer) SetSiteAccess(v string) {
	o.SiteAccess = v
}

// GetEmail returns the Email field value
func (o *Reviewer) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Reviewer) SetEmail(v string) {
	o.Email = v
}

// GetFullName returns the FullName field value
func (o *Reviewer) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *Reviewer) SetFullName(v string) {
	o.FullName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Reviewer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Reviewer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Reviewer) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Reviewer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Reviewer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Reviewer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reviewer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["account_id"] = o.AccountId
	toSerialize["user_id"] = o.UserId
	toSerialize["state"] = o.State
	toSerialize["site_access"] = o.SiteAccess
	toSerialize["email"] = o.Email
	toSerialize["full_name"] = o.FullName
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Reviewer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"account_id",
		"user_id",
		"state",
		"site_access",
		"email",
		"full_name",
		"created_at",
		"updated_at",
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

	varReviewer := _Reviewer{}

	err = json.Unmarshal(data, &varReviewer)

	if err != nil {
		return err
	}

	*o = Reviewer(varReviewer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "site_access")
		delete(additionalProperties, "email")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewer struct {
	value *Reviewer
	isSet bool
}

func (v NullableReviewer) Get() *Reviewer {
	return v.value
}

func (v *NullableReviewer) Set(val *Reviewer) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewer) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewer(val *Reviewer) *NullableReviewer {
	return &NullableReviewer{value: val, isSet: true}
}

func (v NullableReviewer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


