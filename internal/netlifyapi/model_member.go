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

// checks if the Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Member{}

// Member Member model definition. Similar to the User model, but includes the information of the associated team (account).
type Member struct {
	// the identifier for the member (user ID)
	Id string `json:"id"`
	// the full name of the member
	FullName string `json:"full_name"`
	// the email of the member
	Email string `json:"email"`
	// the avatar URL of the member
	Avatar string `json:"avatar"`
	// the role of the member in the team
	Role string `json:"role"`
	// if the member has 2Factor Auth enabled or
	MfaEnabled bool `json:"mfa_enabled"`
	// the table of what CRUD actions that the member could take in the UI
	Capabilities map[string]interface{} `json:"capabilities"`
	ConnectedAccounts CDPUserConnectedAccounts `json:"connected_accounts"`
	// site ids this member has access to
	SiteIds []string `json:"site_ids"`
	// which site access this member has. either all, selected, or none
	SiteAccess string `json:"site_access"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// if the associated invite is pending acceptance
	Pending bool `json:"pending"`
	// flag indicating if the member is managed by an idp directory
	ManagedByDirectorySync bool `json:"managed_by_directory_sync"`
	CommitterMatchMethod *MemberCommitterMatchMethod `json:"committer_match_method,omitempty"`
	// The last activity date of the member on the site. Note: activity data cannot be accessed before Nov 17th 2023, null values indicate last activity was before this date.
	LastActivityDate *string `json:"last_activity_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Member Member

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember(id string, fullName string, email string, avatar string, role string, mfaEnabled bool, capabilities map[string]interface{}, connectedAccounts CDPUserConnectedAccounts, siteIds []string, siteAccess string, createdAt time.Time, updatedAt time.Time, pending bool, managedByDirectorySync bool) *Member {
	this := Member{}
	this.Id = id
	this.FullName = fullName
	this.Email = email
	this.Avatar = avatar
	this.Role = role
	this.MfaEnabled = mfaEnabled
	this.Capabilities = capabilities
	this.ConnectedAccounts = connectedAccounts
	this.SiteIds = siteIds
	this.SiteAccess = siteAccess
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Pending = pending
	this.ManagedByDirectorySync = managedByDirectorySync
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetId returns the Id field value
func (o *Member) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Member) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Member) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *Member) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *Member) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *Member) SetFullName(v string) {
	o.FullName = v
}

// GetEmail returns the Email field value
func (o *Member) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Member) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Member) SetEmail(v string) {
	o.Email = v
}

// GetAvatar returns the Avatar field value
func (o *Member) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *Member) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *Member) SetAvatar(v string) {
	o.Avatar = v
}

// GetRole returns the Role field value
func (o *Member) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Member) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Member) SetRole(v string) {
	o.Role = v
}

// GetMfaEnabled returns the MfaEnabled field value
func (o *Member) GetMfaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MfaEnabled
}

// GetMfaEnabledOk returns a tuple with the MfaEnabled field value
// and a boolean to check if the value has been set.
func (o *Member) GetMfaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaEnabled, true
}

// SetMfaEnabled sets field value
func (o *Member) SetMfaEnabled(v bool) {
	o.MfaEnabled = v
}

// GetCapabilities returns the Capabilities field value
func (o *Member) GetCapabilities() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *Member) GetCapabilitiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *Member) SetCapabilities(v map[string]interface{}) {
	o.Capabilities = v
}

// GetConnectedAccounts returns the ConnectedAccounts field value
func (o *Member) GetConnectedAccounts() CDPUserConnectedAccounts {
	if o == nil {
		var ret CDPUserConnectedAccounts
		return ret
	}

	return o.ConnectedAccounts
}

// GetConnectedAccountsOk returns a tuple with the ConnectedAccounts field value
// and a boolean to check if the value has been set.
func (o *Member) GetConnectedAccountsOk() (*CDPUserConnectedAccounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedAccounts, true
}

// SetConnectedAccounts sets field value
func (o *Member) SetConnectedAccounts(v CDPUserConnectedAccounts) {
	o.ConnectedAccounts = v
}

// GetSiteIds returns the SiteIds field value
func (o *Member) GetSiteIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SiteIds
}

// GetSiteIdsOk returns a tuple with the SiteIds field value
// and a boolean to check if the value has been set.
func (o *Member) GetSiteIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SiteIds, true
}

// SetSiteIds sets field value
func (o *Member) SetSiteIds(v []string) {
	o.SiteIds = v
}

// GetSiteAccess returns the SiteAccess field value
func (o *Member) GetSiteAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteAccess
}

// GetSiteAccessOk returns a tuple with the SiteAccess field value
// and a boolean to check if the value has been set.
func (o *Member) GetSiteAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteAccess, true
}

// SetSiteAccess sets field value
func (o *Member) SetSiteAccess(v string) {
	o.SiteAccess = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Member) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Member) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Member) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Member) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Member) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Member) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetPending returns the Pending field value
func (o *Member) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *Member) GetPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *Member) SetPending(v bool) {
	o.Pending = v
}

// GetManagedByDirectorySync returns the ManagedByDirectorySync field value
func (o *Member) GetManagedByDirectorySync() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ManagedByDirectorySync
}

// GetManagedByDirectorySyncOk returns a tuple with the ManagedByDirectorySync field value
// and a boolean to check if the value has been set.
func (o *Member) GetManagedByDirectorySyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagedByDirectorySync, true
}

// SetManagedByDirectorySync sets field value
func (o *Member) SetManagedByDirectorySync(v bool) {
	o.ManagedByDirectorySync = v
}

// GetCommitterMatchMethod returns the CommitterMatchMethod field value if set, zero value otherwise.
func (o *Member) GetCommitterMatchMethod() MemberCommitterMatchMethod {
	if o == nil || IsNil(o.CommitterMatchMethod) {
		var ret MemberCommitterMatchMethod
		return ret
	}
	return *o.CommitterMatchMethod
}

// GetCommitterMatchMethodOk returns a tuple with the CommitterMatchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetCommitterMatchMethodOk() (*MemberCommitterMatchMethod, bool) {
	if o == nil || IsNil(o.CommitterMatchMethod) {
		return nil, false
	}
	return o.CommitterMatchMethod, true
}

// HasCommitterMatchMethod returns a boolean if a field has been set.
func (o *Member) HasCommitterMatchMethod() bool {
	if o != nil && !IsNil(o.CommitterMatchMethod) {
		return true
	}

	return false
}

// SetCommitterMatchMethod gets a reference to the given MemberCommitterMatchMethod and assigns it to the CommitterMatchMethod field.
func (o *Member) SetCommitterMatchMethod(v MemberCommitterMatchMethod) {
	o.CommitterMatchMethod = &v
}

// GetLastActivityDate returns the LastActivityDate field value if set, zero value otherwise.
func (o *Member) GetLastActivityDate() string {
	if o == nil || IsNil(o.LastActivityDate) {
		var ret string
		return ret
	}
	return *o.LastActivityDate
}

// GetLastActivityDateOk returns a tuple with the LastActivityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetLastActivityDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastActivityDate) {
		return nil, false
	}
	return o.LastActivityDate, true
}

// HasLastActivityDate returns a boolean if a field has been set.
func (o *Member) HasLastActivityDate() bool {
	if o != nil && !IsNil(o.LastActivityDate) {
		return true
	}

	return false
}

// SetLastActivityDate gets a reference to the given string and assigns it to the LastActivityDate field.
func (o *Member) SetLastActivityDate(v string) {
	o.LastActivityDate = &v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["full_name"] = o.FullName
	toSerialize["email"] = o.Email
	toSerialize["avatar"] = o.Avatar
	toSerialize["role"] = o.Role
	toSerialize["mfa_enabled"] = o.MfaEnabled
	toSerialize["capabilities"] = o.Capabilities
	toSerialize["connected_accounts"] = o.ConnectedAccounts
	toSerialize["site_ids"] = o.SiteIds
	toSerialize["site_access"] = o.SiteAccess
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["pending"] = o.Pending
	toSerialize["managed_by_directory_sync"] = o.ManagedByDirectorySync
	if !IsNil(o.CommitterMatchMethod) {
		toSerialize["committer_match_method"] = o.CommitterMatchMethod
	}
	if !IsNil(o.LastActivityDate) {
		toSerialize["last_activity_date"] = o.LastActivityDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Member) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"full_name",
		"email",
		"avatar",
		"role",
		"mfa_enabled",
		"capabilities",
		"connected_accounts",
		"site_ids",
		"site_access",
		"created_at",
		"updated_at",
		"pending",
		"managed_by_directory_sync",
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

	varMember := _Member{}

	err = json.Unmarshal(data, &varMember)

	if err != nil {
		return err
	}

	*o = Member(varMember)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "avatar")
		delete(additionalProperties, "role")
		delete(additionalProperties, "mfa_enabled")
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "connected_accounts")
		delete(additionalProperties, "site_ids")
		delete(additionalProperties, "site_access")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "managed_by_directory_sync")
		delete(additionalProperties, "committer_match_method")
		delete(additionalProperties, "last_activity_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


