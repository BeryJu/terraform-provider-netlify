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

// checks if the DeployedBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployedBranch{}

// DeployedBranch DeployedBranch model definition
type DeployedBranch struct {
	// The ID of the deployed branch
	Id string `json:"id"`
	// The site ID of the deployed branch
	SiteId string `json:"site_id"`
	// The deploy ID of the deployed branch
	DeployId string `json:"deploy_id"`
	// The name of the deployed branch
	Name string `json:"name"`
	// The slugified name of the deployed branch
	Slug string `json:"slug"`
	// The URL of the deployed branch
	Url string `json:"url"`
	// The SSL URL of the deployed branch
	SslUrl string `json:"ssl_url"`
	// When the deployed branch was created
	CreatedAt time.Time `json:"created_at"`
	// When the deployed branch was updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _DeployedBranch DeployedBranch

// NewDeployedBranch instantiates a new DeployedBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployedBranch(id string, siteId string, deployId string, name string, slug string, url string, sslUrl string, createdAt time.Time, updatedAt time.Time) *DeployedBranch {
	this := DeployedBranch{}
	this.Id = id
	this.SiteId = siteId
	this.DeployId = deployId
	this.Name = name
	this.Slug = slug
	this.Url = url
	this.SslUrl = sslUrl
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDeployedBranchWithDefaults instantiates a new DeployedBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployedBranchWithDefaults() *DeployedBranch {
	this := DeployedBranch{}
	return &this
}

// GetId returns the Id field value
func (o *DeployedBranch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeployedBranch) SetId(v string) {
	o.Id = v
}

// GetSiteId returns the SiteId field value
func (o *DeployedBranch) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *DeployedBranch) SetSiteId(v string) {
	o.SiteId = v
}

// GetDeployId returns the DeployId field value
func (o *DeployedBranch) GetDeployId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployId
}

// GetDeployIdOk returns a tuple with the DeployId field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetDeployIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployId, true
}

// SetDeployId sets field value
func (o *DeployedBranch) SetDeployId(v string) {
	o.DeployId = v
}

// GetName returns the Name field value
func (o *DeployedBranch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeployedBranch) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *DeployedBranch) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DeployedBranch) SetSlug(v string) {
	o.Slug = v
}

// GetUrl returns the Url field value
func (o *DeployedBranch) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeployedBranch) SetUrl(v string) {
	o.Url = v
}

// GetSslUrl returns the SslUrl field value
func (o *DeployedBranch) GetSslUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SslUrl
}

// GetSslUrlOk returns a tuple with the SslUrl field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetSslUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SslUrl, true
}

// SetSslUrl sets field value
func (o *DeployedBranch) SetSslUrl(v string) {
	o.SslUrl = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeployedBranch) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeployedBranch) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DeployedBranch) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DeployedBranch) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DeployedBranch) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o DeployedBranch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployedBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["site_id"] = o.SiteId
	toSerialize["deploy_id"] = o.DeployId
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["url"] = o.Url
	toSerialize["ssl_url"] = o.SslUrl
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeployedBranch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"site_id",
		"deploy_id",
		"name",
		"slug",
		"url",
		"ssl_url",
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

	varDeployedBranch := _DeployedBranch{}

	err = json.Unmarshal(data, &varDeployedBranch)

	if err != nil {
		return err
	}

	*o = DeployedBranch(varDeployedBranch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "deploy_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "url")
		delete(additionalProperties, "ssl_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeployedBranch struct {
	value *DeployedBranch
	isSet bool
}

func (v NullableDeployedBranch) Get() *DeployedBranch {
	return v.value
}

func (v *NullableDeployedBranch) Set(val *DeployedBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployedBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployedBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployedBranch(val *DeployedBranch) *NullableDeployedBranch {
	return &NullableDeployedBranch{value: val, isSet: true}
}

func (v NullableDeployedBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployedBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


