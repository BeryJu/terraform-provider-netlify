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

// TrafficRuleActionConfig - struct for TrafficRuleActionConfig
type TrafficRuleActionConfig struct {
	ResponseConfig *ResponseConfig
	RewriteConfig *RewriteConfig
}

// ResponseConfigAsTrafficRuleActionConfig is a convenience function that returns ResponseConfig wrapped in TrafficRuleActionConfig
func ResponseConfigAsTrafficRuleActionConfig(v *ResponseConfig) TrafficRuleActionConfig {
	return TrafficRuleActionConfig{
		ResponseConfig: v,
	}
}

// RewriteConfigAsTrafficRuleActionConfig is a convenience function that returns RewriteConfig wrapped in TrafficRuleActionConfig
func RewriteConfigAsTrafficRuleActionConfig(v *RewriteConfig) TrafficRuleActionConfig {
	return TrafficRuleActionConfig{
		RewriteConfig: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TrafficRuleActionConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ResponseConfig
	err = newStrictDecoder(data).Decode(&dst.ResponseConfig)
	if err == nil {
		jsonResponseConfig, _ := json.Marshal(dst.ResponseConfig)
		if string(jsonResponseConfig) == "{}" { // empty struct
			dst.ResponseConfig = nil
		} else {
			match++
		}
	} else {
		dst.ResponseConfig = nil
	}

	// try to unmarshal data into RewriteConfig
	err = newStrictDecoder(data).Decode(&dst.RewriteConfig)
	if err == nil {
		jsonRewriteConfig, _ := json.Marshal(dst.RewriteConfig)
		if string(jsonRewriteConfig) == "{}" { // empty struct
			dst.RewriteConfig = nil
		} else {
			match++
		}
	} else {
		dst.RewriteConfig = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ResponseConfig = nil
		dst.RewriteConfig = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TrafficRuleActionConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TrafficRuleActionConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TrafficRuleActionConfig) MarshalJSON() ([]byte, error) {
	if src.ResponseConfig != nil {
		return json.Marshal(&src.ResponseConfig)
	}

	if src.RewriteConfig != nil {
		return json.Marshal(&src.RewriteConfig)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TrafficRuleActionConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ResponseConfig != nil {
		return obj.ResponseConfig
	}

	if obj.RewriteConfig != nil {
		return obj.RewriteConfig
	}

	// all schemas are nil
	return nil
}

type NullableTrafficRuleActionConfig struct {
	value *TrafficRuleActionConfig
	isSet bool
}

func (v NullableTrafficRuleActionConfig) Get() *TrafficRuleActionConfig {
	return v.value
}

func (v *NullableTrafficRuleActionConfig) Set(val *TrafficRuleActionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficRuleActionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficRuleActionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficRuleActionConfig(val *TrafficRuleActionConfig) *NullableTrafficRuleActionConfig {
	return &NullableTrafficRuleActionConfig{value: val, isSet: true}
}

func (v NullableTrafficRuleActionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficRuleActionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

