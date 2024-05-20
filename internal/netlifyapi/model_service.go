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

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	ServicePath string `json:"service_path"`
	LongDescription string `json:"long_description"`
	Description string `json:"description"`
	Events []map[string]interface{} `json:"events"`
	Tags []string `json:"tags"`
	Icon string `json:"icon"`
	ManifestUrl string `json:"manifest_url"`
	Environments []string `json:"environments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Service Service

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(id string, name string, slug string, servicePath string, longDescription string, description string, events []map[string]interface{}, tags []string, icon string, manifestUrl string, environments []string, createdAt time.Time, updatedAt time.Time) *Service {
	this := Service{}
	this.Id = id
	this.Name = name
	this.Slug = slug
	this.ServicePath = servicePath
	this.LongDescription = longDescription
	this.Description = description
	this.Events = events
	this.Tags = tags
	this.Icon = icon
	this.ManifestUrl = manifestUrl
	this.Environments = environments
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value
func (o *Service) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Service) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Service) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Service) SetSlug(v string) {
	o.Slug = v
}

// GetServicePath returns the ServicePath field value
func (o *Service) GetServicePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServicePath
}

// GetServicePathOk returns a tuple with the ServicePath field value
// and a boolean to check if the value has been set.
func (o *Service) GetServicePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServicePath, true
}

// SetServicePath sets field value
func (o *Service) SetServicePath(v string) {
	o.ServicePath = v
}

// GetLongDescription returns the LongDescription field value
func (o *Service) GetLongDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LongDescription
}

// GetLongDescriptionOk returns a tuple with the LongDescription field value
// and a boolean to check if the value has been set.
func (o *Service) GetLongDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongDescription, true
}

// SetLongDescription sets field value
func (o *Service) SetLongDescription(v string) {
	o.LongDescription = v
}

// GetDescription returns the Description field value
func (o *Service) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Service) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Service) SetDescription(v string) {
	o.Description = v
}

// GetEvents returns the Events field value
func (o *Service) GetEvents() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Service) GetEventsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *Service) SetEvents(v []map[string]interface{}) {
	o.Events = v
}

// GetTags returns the Tags field value
func (o *Service) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Service) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Service) SetTags(v []string) {
	o.Tags = v
}

// GetIcon returns the Icon field value
func (o *Service) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *Service) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *Service) SetIcon(v string) {
	o.Icon = v
}

// GetManifestUrl returns the ManifestUrl field value
func (o *Service) GetManifestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManifestUrl
}

// GetManifestUrlOk returns a tuple with the ManifestUrl field value
// and a boolean to check if the value has been set.
func (o *Service) GetManifestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManifestUrl, true
}

// SetManifestUrl sets field value
func (o *Service) SetManifestUrl(v string) {
	o.ManifestUrl = v
}

// GetEnvironments returns the Environments field value
func (o *Service) GetEnvironments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *Service) GetEnvironmentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environments, true
}

// SetEnvironments sets field value
func (o *Service) SetEnvironments(v []string) {
	o.Environments = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Service) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Service) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Service) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Service) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Service) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["service_path"] = o.ServicePath
	toSerialize["long_description"] = o.LongDescription
	toSerialize["description"] = o.Description
	toSerialize["events"] = o.Events
	toSerialize["tags"] = o.Tags
	toSerialize["icon"] = o.Icon
	toSerialize["manifest_url"] = o.ManifestUrl
	toSerialize["environments"] = o.Environments
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Service) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"slug",
		"service_path",
		"long_description",
		"description",
		"events",
		"tags",
		"icon",
		"manifest_url",
		"environments",
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

	varService := _Service{}

	err = json.Unmarshal(data, &varService)

	if err != nil {
		return err
	}

	*o = Service(varService)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "service_path")
		delete(additionalProperties, "long_description")
		delete(additionalProperties, "description")
		delete(additionalProperties, "events")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "manifest_url")
		delete(additionalProperties, "environments")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


