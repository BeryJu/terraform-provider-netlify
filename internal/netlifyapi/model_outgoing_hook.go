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

// checks if the OutgoingHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutgoingHook{}

// OutgoingHook Outgoing hook model definition
type OutgoingHook struct {
	// The ID of the hook
	Id string `json:"id"`
	// The site ID associated with this hook
	SiteId string `json:"site_id"`
	// The form ID associated with this hook
	FormId string `json:"form_id"`
	// The form name of this hook
	FormName string `json:"form_name"`
	// The user ID of the hook creator
	UserId string `json:"user_id"`
	// The type of the hook
	Type string `json:"type"`
	// The name of the hook event
	Event string `json:"event"`
	// The additional data for the hook
	Data map[string]interface{} `json:"data"`
	// Whether the last notification succeeded
	Success bool `json:"success"`
	// When the hook was created
	CreatedAt time.Time `json:"created_at"`
	// When the hook was updated
	UpdatedAt time.Time `json:"updated_at"`
	// Which actor the hook is triggered by
	Actor string `json:"actor"`
	// Whether the hook is disabled due to the error rate
	Disabled bool `json:"disabled"`
	// Whether this hook type is supported by the account type
	Restricted bool `json:"restricted"`
	AdditionalProperties map[string]interface{}
}

type _OutgoingHook OutgoingHook

// NewOutgoingHook instantiates a new OutgoingHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingHook(id string, siteId string, formId string, formName string, userId string, type_ string, event string, data map[string]interface{}, success bool, createdAt time.Time, updatedAt time.Time, actor string, disabled bool, restricted bool) *OutgoingHook {
	this := OutgoingHook{}
	this.Id = id
	this.SiteId = siteId
	this.FormId = formId
	this.FormName = formName
	this.UserId = userId
	this.Type = type_
	this.Event = event
	this.Data = data
	this.Success = success
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Actor = actor
	this.Disabled = disabled
	this.Restricted = restricted
	return &this
}

// NewOutgoingHookWithDefaults instantiates a new OutgoingHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingHookWithDefaults() *OutgoingHook {
	this := OutgoingHook{}
	return &this
}

// GetId returns the Id field value
func (o *OutgoingHook) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OutgoingHook) SetId(v string) {
	o.Id = v
}

// GetSiteId returns the SiteId field value
func (o *OutgoingHook) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *OutgoingHook) SetSiteId(v string) {
	o.SiteId = v
}

// GetFormId returns the FormId field value
func (o *OutgoingHook) GetFormId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetFormIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *OutgoingHook) SetFormId(v string) {
	o.FormId = v
}

// GetFormName returns the FormName field value
func (o *OutgoingHook) GetFormName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetFormNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormName, true
}

// SetFormName sets field value
func (o *OutgoingHook) SetFormName(v string) {
	o.FormName = v
}

// GetUserId returns the UserId field value
func (o *OutgoingHook) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *OutgoingHook) SetUserId(v string) {
	o.UserId = v
}

// GetType returns the Type field value
func (o *OutgoingHook) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OutgoingHook) SetType(v string) {
	o.Type = v
}

// GetEvent returns the Event field value
func (o *OutgoingHook) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *OutgoingHook) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *OutgoingHook) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *OutgoingHook) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetSuccess returns the Success field value
func (o *OutgoingHook) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *OutgoingHook) SetSuccess(v bool) {
	o.Success = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OutgoingHook) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OutgoingHook) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OutgoingHook) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OutgoingHook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetActor returns the Actor field value
func (o *OutgoingHook) GetActor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetActorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *OutgoingHook) SetActor(v string) {
	o.Actor = v
}

// GetDisabled returns the Disabled field value
func (o *OutgoingHook) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *OutgoingHook) SetDisabled(v bool) {
	o.Disabled = v
}

// GetRestricted returns the Restricted field value
func (o *OutgoingHook) GetRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value
// and a boolean to check if the value has been set.
func (o *OutgoingHook) GetRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Restricted, true
}

// SetRestricted sets field value
func (o *OutgoingHook) SetRestricted(v bool) {
	o.Restricted = v
}

func (o OutgoingHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutgoingHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["site_id"] = o.SiteId
	toSerialize["form_id"] = o.FormId
	toSerialize["form_name"] = o.FormName
	toSerialize["user_id"] = o.UserId
	toSerialize["type"] = o.Type
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	toSerialize["success"] = o.Success
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["actor"] = o.Actor
	toSerialize["disabled"] = o.Disabled
	toSerialize["restricted"] = o.Restricted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutgoingHook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"site_id",
		"form_id",
		"form_name",
		"user_id",
		"type",
		"event",
		"data",
		"success",
		"created_at",
		"updated_at",
		"actor",
		"disabled",
		"restricted",
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

	varOutgoingHook := _OutgoingHook{}

	err = json.Unmarshal(data, &varOutgoingHook)

	if err != nil {
		return err
	}

	*o = OutgoingHook(varOutgoingHook)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "form_id")
		delete(additionalProperties, "form_name")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		delete(additionalProperties, "success")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "restricted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutgoingHook struct {
	value *OutgoingHook
	isSet bool
}

func (v NullableOutgoingHook) Get() *OutgoingHook {
	return v.value
}

func (v *NullableOutgoingHook) Set(val *OutgoingHook) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingHook) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingHook(val *OutgoingHook) *NullableOutgoingHook {
	return &NullableOutgoingHook{value: val, isSet: true}
}

func (v NullableOutgoingHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


