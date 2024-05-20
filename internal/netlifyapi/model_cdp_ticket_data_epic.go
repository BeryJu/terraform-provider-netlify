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

// checks if the CDPTicketDataEpic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CDPTicketDataEpic{}

// CDPTicketDataEpic struct for CDPTicketDataEpic
type CDPTicketDataEpic struct {
	Archived bool `json:"archived"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description"`
	EndedAt time.Time `json:"ended_at"`
	Id string `json:"id"`
	Name string `json:"name"`
	StateType string `json:"state_type"`
	StartedAt time.Time `json:"started_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CDPTicketDataEpic CDPTicketDataEpic

// NewCDPTicketDataEpic instantiates a new CDPTicketDataEpic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCDPTicketDataEpic(archived bool, completed bool, createdAt time.Time, description string, endedAt time.Time, id string, name string, stateType string, startedAt time.Time, updatedAt time.Time) *CDPTicketDataEpic {
	this := CDPTicketDataEpic{}
	this.Archived = archived
	this.Completed = completed
	this.CreatedAt = createdAt
	this.Description = description
	this.EndedAt = endedAt
	this.Id = id
	this.Name = name
	this.StateType = stateType
	this.StartedAt = startedAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCDPTicketDataEpicWithDefaults instantiates a new CDPTicketDataEpic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCDPTicketDataEpicWithDefaults() *CDPTicketDataEpic {
	this := CDPTicketDataEpic{}
	return &this
}

// GetArchived returns the Archived field value
func (o *CDPTicketDataEpic) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *CDPTicketDataEpic) SetArchived(v bool) {
	o.Archived = v
}

// GetCompleted returns the Completed field value
func (o *CDPTicketDataEpic) GetCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *CDPTicketDataEpic) SetCompleted(v bool) {
	o.Completed = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CDPTicketDataEpic) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CDPTicketDataEpic) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *CDPTicketDataEpic) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CDPTicketDataEpic) SetDescription(v string) {
	o.Description = v
}

// GetEndedAt returns the EndedAt field value
func (o *CDPTicketDataEpic) GetEndedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndedAt, true
}

// SetEndedAt sets field value
func (o *CDPTicketDataEpic) SetEndedAt(v time.Time) {
	o.EndedAt = v
}

// GetId returns the Id field value
func (o *CDPTicketDataEpic) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CDPTicketDataEpic) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CDPTicketDataEpic) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CDPTicketDataEpic) SetName(v string) {
	o.Name = v
}

// GetStateType returns the StateType field value
func (o *CDPTicketDataEpic) GetStateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateType
}

// GetStateTypeOk returns a tuple with the StateType field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetStateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateType, true
}

// SetStateType sets field value
func (o *CDPTicketDataEpic) SetStateType(v string) {
	o.StateType = v
}

// GetStartedAt returns the StartedAt field value
func (o *CDPTicketDataEpic) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *CDPTicketDataEpic) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CDPTicketDataEpic) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataEpic) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CDPTicketDataEpic) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CDPTicketDataEpic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CDPTicketDataEpic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archived"] = o.Archived
	toSerialize["completed"] = o.Completed
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["ended_at"] = o.EndedAt
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["state_type"] = o.StateType
	toSerialize["started_at"] = o.StartedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CDPTicketDataEpic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archived",
		"completed",
		"created_at",
		"description",
		"ended_at",
		"id",
		"name",
		"state_type",
		"started_at",
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

	varCDPTicketDataEpic := _CDPTicketDataEpic{}

	err = json.Unmarshal(data, &varCDPTicketDataEpic)

	if err != nil {
		return err
	}

	*o = CDPTicketDataEpic(varCDPTicketDataEpic)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "archived")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ended_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state_type")
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCDPTicketDataEpic struct {
	value *CDPTicketDataEpic
	isSet bool
}

func (v NullableCDPTicketDataEpic) Get() *CDPTicketDataEpic {
	return v.value
}

func (v *NullableCDPTicketDataEpic) Set(val *CDPTicketDataEpic) {
	v.value = val
	v.isSet = true
}

func (v NullableCDPTicketDataEpic) IsSet() bool {
	return v.isSet
}

func (v *NullableCDPTicketDataEpic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCDPTicketDataEpic(val *CDPTicketDataEpic) *NullableCDPTicketDataEpic {
	return &NullableCDPTicketDataEpic{value: val, isSet: true}
}

func (v NullableCDPTicketDataEpic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCDPTicketDataEpic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


