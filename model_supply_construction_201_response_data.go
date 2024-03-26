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

// checks if the SupplyConstruction201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyConstruction201ResponseData{}

// SupplyConstruction201ResponseData struct for SupplyConstruction201ResponseData
type SupplyConstruction201ResponseData struct {
	Construction Construction `json:"construction"`
	Cargo ShipCargo `json:"cargo"`
}

type _SupplyConstruction201ResponseData SupplyConstruction201ResponseData

// NewSupplyConstruction201ResponseData instantiates a new SupplyConstruction201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyConstruction201ResponseData(construction Construction, cargo ShipCargo) *SupplyConstruction201ResponseData {
	this := SupplyConstruction201ResponseData{}
	this.Construction = construction
	this.Cargo = cargo
	return &this
}

// NewSupplyConstruction201ResponseDataWithDefaults instantiates a new SupplyConstruction201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyConstruction201ResponseDataWithDefaults() *SupplyConstruction201ResponseData {
	this := SupplyConstruction201ResponseData{}
	return &this
}

// GetConstruction returns the Construction field value
func (o *SupplyConstruction201ResponseData) GetConstruction() Construction {
	if o == nil {
		var ret Construction
		return ret
	}

	return o.Construction
}

// GetConstructionOk returns a tuple with the Construction field value
// and a boolean to check if the value has been set.
func (o *SupplyConstruction201ResponseData) GetConstructionOk() (*Construction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Construction, true
}

// SetConstruction sets field value
func (o *SupplyConstruction201ResponseData) SetConstruction(v Construction) {
	o.Construction = v
}

// GetCargo returns the Cargo field value
func (o *SupplyConstruction201ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *SupplyConstruction201ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *SupplyConstruction201ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

func (o SupplyConstruction201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyConstruction201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["construction"] = o.Construction
	toSerialize["cargo"] = o.Cargo
	return toSerialize, nil
}

func (o *SupplyConstruction201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"construction",
		"cargo",
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

	varSupplyConstruction201ResponseData := _SupplyConstruction201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyConstruction201ResponseData)

	if err != nil {
		return err
	}

	*o = SupplyConstruction201ResponseData(varSupplyConstruction201ResponseData)

	return err
}

type NullableSupplyConstruction201ResponseData struct {
	value *SupplyConstruction201ResponseData
	isSet bool
}

func (v NullableSupplyConstruction201ResponseData) Get() *SupplyConstruction201ResponseData {
	return v.value
}

func (v *NullableSupplyConstruction201ResponseData) Set(val *SupplyConstruction201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyConstruction201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyConstruction201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyConstruction201ResponseData(val *SupplyConstruction201ResponseData) *NullableSupplyConstruction201ResponseData {
	return &NullableSupplyConstruction201ResponseData{value: val, isSet: true}
}

func (v NullableSupplyConstruction201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyConstruction201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


