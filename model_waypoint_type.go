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

// WaypointType The type of waypoint.
type WaypointType string

// List of WaypointType
const (
	WAYPOINTTYPE_PLANET WaypointType = "PLANET"
	WAYPOINTTYPE_GAS_GIANT WaypointType = "GAS_GIANT"
	WAYPOINTTYPE_MOON WaypointType = "MOON"
	WAYPOINTTYPE_ORBITAL_STATION WaypointType = "ORBITAL_STATION"
	WAYPOINTTYPE_JUMP_GATE WaypointType = "JUMP_GATE"
	WAYPOINTTYPE_ASTEROID_FIELD WaypointType = "ASTEROID_FIELD"
	WAYPOINTTYPE_ASTEROID WaypointType = "ASTEROID"
	WAYPOINTTYPE_ENGINEERED_ASTEROID WaypointType = "ENGINEERED_ASTEROID"
	WAYPOINTTYPE_ASTEROID_BASE WaypointType = "ASTEROID_BASE"
	WAYPOINTTYPE_NEBULA WaypointType = "NEBULA"
	WAYPOINTTYPE_DEBRIS_FIELD WaypointType = "DEBRIS_FIELD"
	WAYPOINTTYPE_GRAVITY_WELL WaypointType = "GRAVITY_WELL"
	WAYPOINTTYPE_ARTIFICIAL_GRAVITY_WELL WaypointType = "ARTIFICIAL_GRAVITY_WELL"
	WAYPOINTTYPE_FUEL_STATION WaypointType = "FUEL_STATION"
)

// All allowed values of WaypointType enum
var AllowedWaypointTypeEnumValues = []WaypointType{
	"PLANET",
	"GAS_GIANT",
	"MOON",
	"ORBITAL_STATION",
	"JUMP_GATE",
	"ASTEROID_FIELD",
	"ASTEROID",
	"ENGINEERED_ASTEROID",
	"ASTEROID_BASE",
	"NEBULA",
	"DEBRIS_FIELD",
	"GRAVITY_WELL",
	"ARTIFICIAL_GRAVITY_WELL",
	"FUEL_STATION",
}

func (v *WaypointType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WaypointType(value)
	for _, existing := range AllowedWaypointTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WaypointType", value)
}

// NewWaypointTypeFromValue returns a pointer to a valid WaypointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWaypointTypeFromValue(v string) (*WaypointType, error) {
	ev := WaypointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WaypointType: valid values are %v", v, AllowedWaypointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WaypointType) IsValid() bool {
	for _, existing := range AllowedWaypointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WaypointType value
func (v WaypointType) Ptr() *WaypointType {
	return &v
}

type NullableWaypointType struct {
	value *WaypointType
	isSet bool
}

func (v NullableWaypointType) Get() *WaypointType {
	return v.value
}

func (v *NullableWaypointType) Set(val *WaypointType) {
	v.value = val
	v.isSet = true
}

func (v NullableWaypointType) IsSet() bool {
	return v.isSet
}

func (v *NullableWaypointType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaypointType(val *WaypointType) *NullableWaypointType {
	return &NullableWaypointType{value: val, isSet: true}
}

func (v NullableWaypointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaypointType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

