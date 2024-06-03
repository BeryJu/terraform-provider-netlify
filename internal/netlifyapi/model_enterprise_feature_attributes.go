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

// checks if the EnterpriseFeatureAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseFeatureAttributes{}

// EnterpriseFeatureAttributes EnterpriseFeatureAttributes model definition.
type EnterpriseFeatureAttributes struct {
	// The custom price for the feature in the contract; per unit for features that start with `additional_`
	Price *string `json:"price,omitempty"`
	// The amount included in the contract; consult Account#capabilities first
	IncludedQuantity *int64 `json:"included_quantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnterpriseFeatureAttributes EnterpriseFeatureAttributes

// NewEnterpriseFeatureAttributes instantiates a new EnterpriseFeatureAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseFeatureAttributes() *EnterpriseFeatureAttributes {
	this := EnterpriseFeatureAttributes{}
	return &this
}

// NewEnterpriseFeatureAttributesWithDefaults instantiates a new EnterpriseFeatureAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseFeatureAttributesWithDefaults() *EnterpriseFeatureAttributes {
	this := EnterpriseFeatureAttributes{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *EnterpriseFeatureAttributes) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseFeatureAttributes) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *EnterpriseFeatureAttributes) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *EnterpriseFeatureAttributes) SetPrice(v string) {
	o.Price = &v
}

// GetIncludedQuantity returns the IncludedQuantity field value if set, zero value otherwise.
func (o *EnterpriseFeatureAttributes) GetIncludedQuantity() int64 {
	if o == nil || IsNil(o.IncludedQuantity) {
		var ret int64
		return ret
	}
	return *o.IncludedQuantity
}

// GetIncludedQuantityOk returns a tuple with the IncludedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseFeatureAttributes) GetIncludedQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.IncludedQuantity) {
		return nil, false
	}
	return o.IncludedQuantity, true
}

// HasIncludedQuantity returns a boolean if a field has been set.
func (o *EnterpriseFeatureAttributes) HasIncludedQuantity() bool {
	if o != nil && !IsNil(o.IncludedQuantity) {
		return true
	}

	return false
}

// SetIncludedQuantity gets a reference to the given int64 and assigns it to the IncludedQuantity field.
func (o *EnterpriseFeatureAttributes) SetIncludedQuantity(v int64) {
	o.IncludedQuantity = &v
}

func (o EnterpriseFeatureAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseFeatureAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.IncludedQuantity) {
		toSerialize["included_quantity"] = o.IncludedQuantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnterpriseFeatureAttributes) UnmarshalJSON(data []byte) (err error) {
	varEnterpriseFeatureAttributes := _EnterpriseFeatureAttributes{}

	err = json.Unmarshal(data, &varEnterpriseFeatureAttributes)

	if err != nil {
		return err
	}

	*o = EnterpriseFeatureAttributes(varEnterpriseFeatureAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "included_quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnterpriseFeatureAttributes struct {
	value *EnterpriseFeatureAttributes
	isSet bool
}

func (v NullableEnterpriseFeatureAttributes) Get() *EnterpriseFeatureAttributes {
	return v.value
}

func (v *NullableEnterpriseFeatureAttributes) Set(val *EnterpriseFeatureAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseFeatureAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseFeatureAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseFeatureAttributes(val *EnterpriseFeatureAttributes) *NullableEnterpriseFeatureAttributes {
	return &NullableEnterpriseFeatureAttributes{value: val, isSet: true}
}

func (v NullableEnterpriseFeatureAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseFeatureAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

