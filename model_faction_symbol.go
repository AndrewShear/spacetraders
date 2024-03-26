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

// FactionSymbol The symbol of the faction.
type FactionSymbol string

// List of FactionSymbol
const (
	FACTIONSYMBOL_COSMIC FactionSymbol = "COSMIC"
	FACTIONSYMBOL_VOID FactionSymbol = "VOID"
	FACTIONSYMBOL_GALACTIC FactionSymbol = "GALACTIC"
	FACTIONSYMBOL_QUANTUM FactionSymbol = "QUANTUM"
	FACTIONSYMBOL_DOMINION FactionSymbol = "DOMINION"
	FACTIONSYMBOL_ASTRO FactionSymbol = "ASTRO"
	FACTIONSYMBOL_CORSAIRS FactionSymbol = "CORSAIRS"
	FACTIONSYMBOL_OBSIDIAN FactionSymbol = "OBSIDIAN"
	FACTIONSYMBOL_AEGIS FactionSymbol = "AEGIS"
	FACTIONSYMBOL_UNITED FactionSymbol = "UNITED"
	FACTIONSYMBOL_SOLITARY FactionSymbol = "SOLITARY"
	FACTIONSYMBOL_COBALT FactionSymbol = "COBALT"
	FACTIONSYMBOL_OMEGA FactionSymbol = "OMEGA"
	FACTIONSYMBOL_ECHO FactionSymbol = "ECHO"
	FACTIONSYMBOL_LORDS FactionSymbol = "LORDS"
	FACTIONSYMBOL_CULT FactionSymbol = "CULT"
	FACTIONSYMBOL_ANCIENTS FactionSymbol = "ANCIENTS"
	FACTIONSYMBOL_SHADOW FactionSymbol = "SHADOW"
	FACTIONSYMBOL_ETHEREAL FactionSymbol = "ETHEREAL"
)

// All allowed values of FactionSymbol enum
var AllowedFactionSymbolEnumValues = []FactionSymbol{
	"COSMIC",
	"VOID",
	"GALACTIC",
	"QUANTUM",
	"DOMINION",
	"ASTRO",
	"CORSAIRS",
	"OBSIDIAN",
	"AEGIS",
	"UNITED",
	"SOLITARY",
	"COBALT",
	"OMEGA",
	"ECHO",
	"LORDS",
	"CULT",
	"ANCIENTS",
	"SHADOW",
	"ETHEREAL",
}

func (v *FactionSymbol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FactionSymbol(value)
	for _, existing := range AllowedFactionSymbolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FactionSymbol", value)
}

// NewFactionSymbolFromValue returns a pointer to a valid FactionSymbol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFactionSymbolFromValue(v string) (*FactionSymbol, error) {
	ev := FactionSymbol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FactionSymbol: valid values are %v", v, AllowedFactionSymbolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FactionSymbol) IsValid() bool {
	for _, existing := range AllowedFactionSymbolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FactionSymbol value
func (v FactionSymbol) Ptr() *FactionSymbol {
	return &v
}

type NullableFactionSymbol struct {
	value *FactionSymbol
	isSet bool
}

func (v NullableFactionSymbol) Get() *FactionSymbol {
	return v.value
}

func (v *NullableFactionSymbol) Set(val *FactionSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableFactionSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableFactionSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactionSymbol(val *FactionSymbol) *NullableFactionSymbol {
	return &NullableFactionSymbol{value: val, isSet: true}
}

func (v NullableFactionSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactionSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

