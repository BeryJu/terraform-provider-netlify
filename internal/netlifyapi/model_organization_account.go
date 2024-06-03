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

// checks if the OrganizationAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAccount{}

// OrganizationAccount Account model definition, with more details relating to organizations.
type OrganizationAccount struct {
	BandwidthUsage BandwidthUsage `json:"bandwidth_usage"`
	// The billing email for the account
	BillingEmail string `json:"billing_email"`
	// The billing name for the account
	BillingName string `json:"billing_name"`
	BuildStatus BuildStatus `json:"build_status"`
	// Whether the HP builds is enabled for the account
	HighPerformanceBuildsEnabled bool `json:"high_performance_builds_enabled"`
	// Whether the HP edge is enabled for the account
	HighPerformanceEdgeEnabled bool `json:"high_performance_edge_enabled"`
	// The identifier for the account (account ID)
	Id string `json:"id"`
	// The number of included concurrent builds
	IncludedConcurrentBuilds int64 `json:"included_concurrent_builds"`
	MemberRoles []map[string]interface{} `json:"member_roles"`
	// The name for the account
	Name string `json:"name"`
	// An array of owner_ids on the account
	OwnerIds []string `json:"owner_ids"`
	// The slug for the account
	Slug string `json:"slug"`
	// The support level of the account
	SupportLevel string `json:"support_level"`
	// Total seats
	TotalSeats int64 `json:"total_seats"`
	// Used seats
	UsedSeats int64 `json:"used_seats"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationAccount OrganizationAccount

// NewOrganizationAccount instantiates a new OrganizationAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAccount(bandwidthUsage BandwidthUsage, billingEmail string, billingName string, buildStatus BuildStatus, highPerformanceBuildsEnabled bool, highPerformanceEdgeEnabled bool, id string, includedConcurrentBuilds int64, memberRoles []map[string]interface{}, name string, ownerIds []string, slug string, supportLevel string, totalSeats int64, usedSeats int64) *OrganizationAccount {
	this := OrganizationAccount{}
	this.BandwidthUsage = bandwidthUsage
	this.BillingEmail = billingEmail
	this.BillingName = billingName
	this.BuildStatus = buildStatus
	this.HighPerformanceBuildsEnabled = highPerformanceBuildsEnabled
	this.HighPerformanceEdgeEnabled = highPerformanceEdgeEnabled
	this.Id = id
	this.IncludedConcurrentBuilds = includedConcurrentBuilds
	this.MemberRoles = memberRoles
	this.Name = name
	this.OwnerIds = ownerIds
	this.Slug = slug
	this.SupportLevel = supportLevel
	this.TotalSeats = totalSeats
	this.UsedSeats = usedSeats
	return &this
}

// NewOrganizationAccountWithDefaults instantiates a new OrganizationAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAccountWithDefaults() *OrganizationAccount {
	this := OrganizationAccount{}
	return &this
}

// GetBandwidthUsage returns the BandwidthUsage field value
func (o *OrganizationAccount) GetBandwidthUsage() BandwidthUsage {
	if o == nil {
		var ret BandwidthUsage
		return ret
	}

	return o.BandwidthUsage
}

// GetBandwidthUsageOk returns a tuple with the BandwidthUsage field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetBandwidthUsageOk() (*BandwidthUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandwidthUsage, true
}

// SetBandwidthUsage sets field value
func (o *OrganizationAccount) SetBandwidthUsage(v BandwidthUsage) {
	o.BandwidthUsage = v
}

// GetBillingEmail returns the BillingEmail field value
func (o *OrganizationAccount) GetBillingEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingEmail
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingEmail, true
}

// SetBillingEmail sets field value
func (o *OrganizationAccount) SetBillingEmail(v string) {
	o.BillingEmail = v
}

// GetBillingName returns the BillingName field value
func (o *OrganizationAccount) GetBillingName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingName
}

// GetBillingNameOk returns a tuple with the BillingName field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetBillingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingName, true
}

// SetBillingName sets field value
func (o *OrganizationAccount) SetBillingName(v string) {
	o.BillingName = v
}

// GetBuildStatus returns the BuildStatus field value
func (o *OrganizationAccount) GetBuildStatus() BuildStatus {
	if o == nil {
		var ret BuildStatus
		return ret
	}

	return o.BuildStatus
}

// GetBuildStatusOk returns a tuple with the BuildStatus field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetBuildStatusOk() (*BuildStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildStatus, true
}

// SetBuildStatus sets field value
func (o *OrganizationAccount) SetBuildStatus(v BuildStatus) {
	o.BuildStatus = v
}

// GetHighPerformanceBuildsEnabled returns the HighPerformanceBuildsEnabled field value
func (o *OrganizationAccount) GetHighPerformanceBuildsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HighPerformanceBuildsEnabled
}

// GetHighPerformanceBuildsEnabledOk returns a tuple with the HighPerformanceBuildsEnabled field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetHighPerformanceBuildsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HighPerformanceBuildsEnabled, true
}

// SetHighPerformanceBuildsEnabled sets field value
func (o *OrganizationAccount) SetHighPerformanceBuildsEnabled(v bool) {
	o.HighPerformanceBuildsEnabled = v
}

// GetHighPerformanceEdgeEnabled returns the HighPerformanceEdgeEnabled field value
func (o *OrganizationAccount) GetHighPerformanceEdgeEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HighPerformanceEdgeEnabled
}

// GetHighPerformanceEdgeEnabledOk returns a tuple with the HighPerformanceEdgeEnabled field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetHighPerformanceEdgeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HighPerformanceEdgeEnabled, true
}

// SetHighPerformanceEdgeEnabled sets field value
func (o *OrganizationAccount) SetHighPerformanceEdgeEnabled(v bool) {
	o.HighPerformanceEdgeEnabled = v
}

// GetId returns the Id field value
func (o *OrganizationAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationAccount) SetId(v string) {
	o.Id = v
}

// GetIncludedConcurrentBuilds returns the IncludedConcurrentBuilds field value
func (o *OrganizationAccount) GetIncludedConcurrentBuilds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IncludedConcurrentBuilds
}

// GetIncludedConcurrentBuildsOk returns a tuple with the IncludedConcurrentBuilds field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetIncludedConcurrentBuildsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludedConcurrentBuilds, true
}

// SetIncludedConcurrentBuilds sets field value
func (o *OrganizationAccount) SetIncludedConcurrentBuilds(v int64) {
	o.IncludedConcurrentBuilds = v
}

// GetMemberRoles returns the MemberRoles field value
func (o *OrganizationAccount) GetMemberRoles() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.MemberRoles
}

// GetMemberRolesOk returns a tuple with the MemberRoles field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetMemberRolesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberRoles, true
}

// SetMemberRoles sets field value
func (o *OrganizationAccount) SetMemberRoles(v []map[string]interface{}) {
	o.MemberRoles = v
}

// GetName returns the Name field value
func (o *OrganizationAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationAccount) SetName(v string) {
	o.Name = v
}

// GetOwnerIds returns the OwnerIds field value
func (o *OrganizationAccount) GetOwnerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OwnerIds
}

// GetOwnerIdsOk returns a tuple with the OwnerIds field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetOwnerIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerIds, true
}

// SetOwnerIds sets field value
func (o *OrganizationAccount) SetOwnerIds(v []string) {
	o.OwnerIds = v
}

// GetSlug returns the Slug field value
func (o *OrganizationAccount) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *OrganizationAccount) SetSlug(v string) {
	o.Slug = v
}

// GetSupportLevel returns the SupportLevel field value
func (o *OrganizationAccount) GetSupportLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportLevel
}

// GetSupportLevelOk returns a tuple with the SupportLevel field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetSupportLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportLevel, true
}

// SetSupportLevel sets field value
func (o *OrganizationAccount) SetSupportLevel(v string) {
	o.SupportLevel = v
}

// GetTotalSeats returns the TotalSeats field value
func (o *OrganizationAccount) GetTotalSeats() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalSeats
}

// GetTotalSeatsOk returns a tuple with the TotalSeats field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetTotalSeatsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSeats, true
}

// SetTotalSeats sets field value
func (o *OrganizationAccount) SetTotalSeats(v int64) {
	o.TotalSeats = v
}

// GetUsedSeats returns the UsedSeats field value
func (o *OrganizationAccount) GetUsedSeats() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UsedSeats
}

// GetUsedSeatsOk returns a tuple with the UsedSeats field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccount) GetUsedSeatsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedSeats, true
}

// SetUsedSeats sets field value
func (o *OrganizationAccount) SetUsedSeats(v int64) {
	o.UsedSeats = v
}

func (o OrganizationAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bandwidth_usage"] = o.BandwidthUsage
	toSerialize["billing_email"] = o.BillingEmail
	toSerialize["billing_name"] = o.BillingName
	toSerialize["build_status"] = o.BuildStatus
	toSerialize["high_performance_builds_enabled"] = o.HighPerformanceBuildsEnabled
	toSerialize["high_performance_edge_enabled"] = o.HighPerformanceEdgeEnabled
	toSerialize["id"] = o.Id
	toSerialize["included_concurrent_builds"] = o.IncludedConcurrentBuilds
	toSerialize["member_roles"] = o.MemberRoles
	toSerialize["name"] = o.Name
	toSerialize["owner_ids"] = o.OwnerIds
	toSerialize["slug"] = o.Slug
	toSerialize["support_level"] = o.SupportLevel
	toSerialize["total_seats"] = o.TotalSeats
	toSerialize["used_seats"] = o.UsedSeats

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bandwidth_usage",
		"billing_email",
		"billing_name",
		"build_status",
		"high_performance_builds_enabled",
		"high_performance_edge_enabled",
		"id",
		"included_concurrent_builds",
		"member_roles",
		"name",
		"owner_ids",
		"slug",
		"support_level",
		"total_seats",
		"used_seats",
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

	varOrganizationAccount := _OrganizationAccount{}

	err = json.Unmarshal(data, &varOrganizationAccount)

	if err != nil {
		return err
	}

	*o = OrganizationAccount(varOrganizationAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth_usage")
		delete(additionalProperties, "billing_email")
		delete(additionalProperties, "billing_name")
		delete(additionalProperties, "build_status")
		delete(additionalProperties, "high_performance_builds_enabled")
		delete(additionalProperties, "high_performance_edge_enabled")
		delete(additionalProperties, "id")
		delete(additionalProperties, "included_concurrent_builds")
		delete(additionalProperties, "member_roles")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_ids")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "support_level")
		delete(additionalProperties, "total_seats")
		delete(additionalProperties, "used_seats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationAccount struct {
	value *OrganizationAccount
	isSet bool
}

func (v NullableOrganizationAccount) Get() *OrganizationAccount {
	return v.value
}

func (v *NullableOrganizationAccount) Set(val *OrganizationAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAccount(val *OrganizationAccount) *NullableOrganizationAccount {
	return &NullableOrganizationAccount{value: val, isSet: true}
}

func (v NullableOrganizationAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

