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

// checks if the GithubEnterpriseConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubEnterpriseConfig{}

// GithubEnterpriseConfig struct for GithubEnterpriseConfig
type GithubEnterpriseConfig struct {
	// The URL for the GitHub instance
	InstanceUrl string `json:"instance_url"`
	// The clone URL for the GitHub instance
	CloneUrl string `json:"clone_url"`
	// The client ID of the GitHub application
	ClientId string `json:"client_id"`
	// The client secret of the GitHub application
	ClientSecret string `json:"client_secret"`
	// The app ID of the GitHub application
	AppId int64 `json:"app_id"`
	// The private key of the GitHub application
	PrivateKey string `json:"private_key"`
	// The webhook secret of the GitHub application
	WebhookSecret string `json:"webhook_secret"`
	AdditionalProperties map[string]interface{}
}

type _GithubEnterpriseConfig GithubEnterpriseConfig

// NewGithubEnterpriseConfig instantiates a new GithubEnterpriseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubEnterpriseConfig(instanceUrl string, cloneUrl string, clientId string, clientSecret string, appId int64, privateKey string, webhookSecret string) *GithubEnterpriseConfig {
	this := GithubEnterpriseConfig{}
	this.InstanceUrl = instanceUrl
	this.CloneUrl = cloneUrl
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.AppId = appId
	this.PrivateKey = privateKey
	this.WebhookSecret = webhookSecret
	return &this
}

// NewGithubEnterpriseConfigWithDefaults instantiates a new GithubEnterpriseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubEnterpriseConfigWithDefaults() *GithubEnterpriseConfig {
	this := GithubEnterpriseConfig{}
	return &this
}

// GetInstanceUrl returns the InstanceUrl field value
func (o *GithubEnterpriseConfig) GetInstanceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetInstanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceUrl, true
}

// SetInstanceUrl sets field value
func (o *GithubEnterpriseConfig) SetInstanceUrl(v string) {
	o.InstanceUrl = v
}

// GetCloneUrl returns the CloneUrl field value
func (o *GithubEnterpriseConfig) GetCloneUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloneUrl
}

// GetCloneUrlOk returns a tuple with the CloneUrl field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetCloneUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloneUrl, true
}

// SetCloneUrl sets field value
func (o *GithubEnterpriseConfig) SetCloneUrl(v string) {
	o.CloneUrl = v
}

// GetClientId returns the ClientId field value
func (o *GithubEnterpriseConfig) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *GithubEnterpriseConfig) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *GithubEnterpriseConfig) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *GithubEnterpriseConfig) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetAppId returns the AppId field value
func (o *GithubEnterpriseConfig) GetAppId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetAppIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *GithubEnterpriseConfig) SetAppId(v int64) {
	o.AppId = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *GithubEnterpriseConfig) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *GithubEnterpriseConfig) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetWebhookSecret returns the WebhookSecret field value
func (o *GithubEnterpriseConfig) GetWebhookSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookSecret
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value
// and a boolean to check if the value has been set.
func (o *GithubEnterpriseConfig) GetWebhookSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookSecret, true
}

// SetWebhookSecret sets field value
func (o *GithubEnterpriseConfig) SetWebhookSecret(v string) {
	o.WebhookSecret = v
}

func (o GithubEnterpriseConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubEnterpriseConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instance_url"] = o.InstanceUrl
	toSerialize["clone_url"] = o.CloneUrl
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["app_id"] = o.AppId
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["webhook_secret"] = o.WebhookSecret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GithubEnterpriseConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instance_url",
		"clone_url",
		"client_id",
		"client_secret",
		"app_id",
		"private_key",
		"webhook_secret",
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

	varGithubEnterpriseConfig := _GithubEnterpriseConfig{}

	err = json.Unmarshal(data, &varGithubEnterpriseConfig)

	if err != nil {
		return err
	}

	*o = GithubEnterpriseConfig(varGithubEnterpriseConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instance_url")
		delete(additionalProperties, "clone_url")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "app_id")
		delete(additionalProperties, "private_key")
		delete(additionalProperties, "webhook_secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGithubEnterpriseConfig struct {
	value *GithubEnterpriseConfig
	isSet bool
}

func (v NullableGithubEnterpriseConfig) Get() *GithubEnterpriseConfig {
	return v.value
}

func (v *NullableGithubEnterpriseConfig) Set(val *GithubEnterpriseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubEnterpriseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubEnterpriseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubEnterpriseConfig(val *GithubEnterpriseConfig) *NullableGithubEnterpriseConfig {
	return &NullableGithubEnterpriseConfig{value: val, isSet: true}
}

func (v NullableGithubEnterpriseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubEnterpriseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


