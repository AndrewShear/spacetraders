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
	"bytes"
	"fmt"
)

// checks if the SiphonYield type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiphonYield{}

// SiphonYield A yield from the siphon operation.
type SiphonYield struct {
	Symbol TradeSymbol `json:"symbol"`
	// The number of units siphoned that were placed into the ship's cargo hold.
	Units int32 `json:"units"`
}

type _SiphonYield SiphonYield

// NewSiphonYield instantiates a new SiphonYield object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiphonYield(symbol TradeSymbol, units int32) *SiphonYield {
	this := SiphonYield{}
	this.Symbol = symbol
	this.Units = units
	return &this
}

// NewSiphonYieldWithDefaults instantiates a new SiphonYield object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiphonYieldWithDefaults() *SiphonYield {
	this := SiphonYield{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SiphonYield) GetSymbol() TradeSymbol {
	if o == nil {
		var ret TradeSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SiphonYield) GetSymbolOk() (*TradeSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SiphonYield) SetSymbol(v TradeSymbol) {
	o.Symbol = v
}

// GetUnits returns the Units field value
func (o *SiphonYield) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *SiphonYield) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *SiphonYield) SetUnits(v int32) {
	o.Units = v
}

func (o SiphonYield) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiphonYield) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["units"] = o.Units
	return toSerialize, nil
}

func (o *SiphonYield) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"units",
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

	varSiphonYield := _SiphonYield{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSiphonYield)

	if err != nil {
		return err
	}

	*o = SiphonYield(varSiphonYield)

	return err
}

type NullableSiphonYield struct {
	value *SiphonYield
	isSet bool
}

func (v NullableSiphonYield) Get() *SiphonYield {
	return v.value
}

func (v *NullableSiphonYield) Set(val *SiphonYield) {
	v.value = val
	v.isSet = true
}

func (v NullableSiphonYield) IsSet() bool {
	return v.isSet
}

func (v *NullableSiphonYield) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiphonYield(val *SiphonYield) *NullableSiphonYield {
	return &NullableSiphonYield{value: val, isSet: true}
}

func (v NullableSiphonYield) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiphonYield) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


