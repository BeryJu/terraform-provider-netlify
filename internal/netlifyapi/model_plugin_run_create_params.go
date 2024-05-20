/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
)

// checks if the PluginRunCreateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginRunCreateParams{}

// PluginRunCreateParams struct for PluginRunCreateParams
type PluginRunCreateParams struct {
	// The name of the plugin
	Package *string `json:"package,omitempty"`
	// The version of the plugin that was run
	Version *string `json:"version,omitempty"`
	// The plugin build lifecycle event that generated the report e.g onPostBuild
	ReportingEvent *string `json:"reporting_event,omitempty"`
	// The resulting success, failure or cancellation state
	State *string `json:"state,omitempty"`
	// The summary reported by the plugin
	Summary *string `json:"summary,omitempty"`
	// The full text result reported by the plugin
	Text *string `json:"text,omitempty"`
	// The ID of the deploy during which the plugin ran
	DeployId *string `json:"deploy_id,omitempty"`
	// The ID of the site for which the plugin ran
	SiteId *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PluginRunCreateParams PluginRunCreateParams

// NewPluginRunCreateParams instantiates a new PluginRunCreateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginRunCreateParams() *PluginRunCreateParams {
	this := PluginRunCreateParams{}
	return &this
}

// NewPluginRunCreateParamsWithDefaults instantiates a new PluginRunCreateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginRunCreateParamsWithDefaults() *PluginRunCreateParams {
	this := PluginRunCreateParams{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetPackage() string {
	if o == nil || IsNil(o.Package) {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetPackageOk() (*string, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *PluginRunCreateParams) SetPackage(v string) {
	o.Package = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PluginRunCreateParams) SetVersion(v string) {
	o.Version = &v
}

// GetReportingEvent returns the ReportingEvent field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetReportingEvent() string {
	if o == nil || IsNil(o.ReportingEvent) {
		var ret string
		return ret
	}
	return *o.ReportingEvent
}

// GetReportingEventOk returns a tuple with the ReportingEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetReportingEventOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingEvent) {
		return nil, false
	}
	return o.ReportingEvent, true
}

// HasReportingEvent returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasReportingEvent() bool {
	if o != nil && !IsNil(o.ReportingEvent) {
		return true
	}

	return false
}

// SetReportingEvent gets a reference to the given string and assigns it to the ReportingEvent field.
func (o *PluginRunCreateParams) SetReportingEvent(v string) {
	o.ReportingEvent = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PluginRunCreateParams) SetState(v string) {
	o.State = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PluginRunCreateParams) SetSummary(v string) {
	o.Summary = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PluginRunCreateParams) SetText(v string) {
	o.Text = &v
}

// GetDeployId returns the DeployId field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetDeployId() string {
	if o == nil || IsNil(o.DeployId) {
		var ret string
		return ret
	}
	return *o.DeployId
}

// GetDeployIdOk returns a tuple with the DeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetDeployIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeployId) {
		return nil, false
	}
	return o.DeployId, true
}

// HasDeployId returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasDeployId() bool {
	if o != nil && !IsNil(o.DeployId) {
		return true
	}

	return false
}

// SetDeployId gets a reference to the given string and assigns it to the DeployId field.
func (o *PluginRunCreateParams) SetDeployId(v string) {
	o.DeployId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *PluginRunCreateParams) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginRunCreateParams) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *PluginRunCreateParams) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *PluginRunCreateParams) SetSiteId(v string) {
	o.SiteId = &v
}

func (o PluginRunCreateParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginRunCreateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ReportingEvent) {
		toSerialize["reporting_event"] = o.ReportingEvent
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.DeployId) {
		toSerialize["deploy_id"] = o.DeployId
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PluginRunCreateParams) UnmarshalJSON(data []byte) (err error) {
	varPluginRunCreateParams := _PluginRunCreateParams{}

	err = json.Unmarshal(data, &varPluginRunCreateParams)

	if err != nil {
		return err
	}

	*o = PluginRunCreateParams(varPluginRunCreateParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "package")
		delete(additionalProperties, "version")
		delete(additionalProperties, "reporting_event")
		delete(additionalProperties, "state")
		delete(additionalProperties, "summary")
		delete(additionalProperties, "text")
		delete(additionalProperties, "deploy_id")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePluginRunCreateParams struct {
	value *PluginRunCreateParams
	isSet bool
}

func (v NullablePluginRunCreateParams) Get() *PluginRunCreateParams {
	return v.value
}

func (v *NullablePluginRunCreateParams) Set(val *PluginRunCreateParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginRunCreateParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginRunCreateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginRunCreateParams(val *PluginRunCreateParams) *NullablePluginRunCreateParams {
	return &NullablePluginRunCreateParams{value: val, isSet: true}
}

func (v NullablePluginRunCreateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginRunCreateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


