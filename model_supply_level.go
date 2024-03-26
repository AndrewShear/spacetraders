/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetraders

import (
	"encoding/json"
	"fmt"
)

// SupplyLevel The supply level of a trade good.
type SupplyLevel string

// List of SupplyLevel
const (
	SUPPLYLEVEL_SCARCE SupplyLevel = "SCARCE"
	SUPPLYLEVEL_LIMITED SupplyLevel = "LIMITED"
	SUPPLYLEVEL_MODERATE SupplyLevel = "MODERATE"
	SUPPLYLEVEL_HIGH SupplyLevel = "HIGH"
	SUPPLYLEVEL_ABUNDANT SupplyLevel = "ABUNDANT"
)

// All allowed values of SupplyLevel enum
var AllowedSupplyLevelEnumValues = []SupplyLevel{
	"SCARCE",
	"LIMITED",
	"MODERATE",
	"HIGH",
	"ABUNDANT",
}

func (v *SupplyLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupplyLevel(value)
	for _, existing := range AllowedSupplyLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupplyLevel", value)
}

// NewSupplyLevelFromValue returns a pointer to a valid SupplyLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupplyLevelFromValue(v string) (*SupplyLevel, error) {
	ev := SupplyLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupplyLevel: valid values are %v", v, AllowedSupplyLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupplyLevel) IsValid() bool {
	for _, existing := range AllowedSupplyLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupplyLevel value
func (v SupplyLevel) Ptr() *SupplyLevel {
	return &v
}

type NullableSupplyLevel struct {
	value *SupplyLevel
	isSet bool
}

func (v NullableSupplyLevel) Get() *SupplyLevel {
	return v.value
}

func (v *NullableSupplyLevel) Set(val *SupplyLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyLevel(val *SupplyLevel) *NullableSupplyLevel {
	return &NullableSupplyLevel{value: val, isSet: true}
}

func (v NullableSupplyLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

