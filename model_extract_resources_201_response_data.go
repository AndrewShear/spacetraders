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

// checks if the ExtractResources201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtractResources201ResponseData{}

// ExtractResources201ResponseData struct for ExtractResources201ResponseData
type ExtractResources201ResponseData struct {
	Cooldown Cooldown `json:"cooldown"`
	Extraction Extraction `json:"extraction"`
	Cargo ShipCargo `json:"cargo"`
	Events []ExtractResources201ResponseDataEventsInner `json:"events"`
}

type _ExtractResources201ResponseData ExtractResources201ResponseData

// NewExtractResources201ResponseData instantiates a new ExtractResources201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractResources201ResponseData(cooldown Cooldown, extraction Extraction, cargo ShipCargo, events []ExtractResources201ResponseDataEventsInner) *ExtractResources201ResponseData {
	this := ExtractResources201ResponseData{}
	this.Cooldown = cooldown
	this.Extraction = extraction
	this.Cargo = cargo
	this.Events = events
	return &this
}

// NewExtractResources201ResponseDataWithDefaults instantiates a new ExtractResources201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractResources201ResponseDataWithDefaults() *ExtractResources201ResponseData {
	this := ExtractResources201ResponseData{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *ExtractResources201ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *ExtractResources201ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *ExtractResources201ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetExtraction returns the Extraction field value
func (o *ExtractResources201ResponseData) GetExtraction() Extraction {
	if o == nil {
		var ret Extraction
		return ret
	}

	return o.Extraction
}

// GetExtractionOk returns a tuple with the Extraction field value
// and a boolean to check if the value has been set.
func (o *ExtractResources201ResponseData) GetExtractionOk() (*Extraction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extraction, true
}

// SetExtraction sets field value
func (o *ExtractResources201ResponseData) SetExtraction(v Extraction) {
	o.Extraction = v
}

// GetCargo returns the Cargo field value
func (o *ExtractResources201ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *ExtractResources201ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *ExtractResources201ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetEvents returns the Events field value
func (o *ExtractResources201ResponseData) GetEvents() []ExtractResources201ResponseDataEventsInner {
	if o == nil {
		var ret []ExtractResources201ResponseDataEventsInner
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ExtractResources201ResponseData) GetEventsOk() ([]ExtractResources201ResponseDataEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ExtractResources201ResponseData) SetEvents(v []ExtractResources201ResponseDataEventsInner) {
	o.Events = v
}

func (o ExtractResources201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtractResources201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["extraction"] = o.Extraction
	toSerialize["cargo"] = o.Cargo
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

func (o *ExtractResources201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
		"extraction",
		"cargo",
		"events",
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

	varExtractResources201ResponseData := _ExtractResources201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtractResources201ResponseData)

	if err != nil {
		return err
	}

	*o = ExtractResources201ResponseData(varExtractResources201ResponseData)

	return err
}

type NullableExtractResources201ResponseData struct {
	value *ExtractResources201ResponseData
	isSet bool
}

func (v NullableExtractResources201ResponseData) Get() *ExtractResources201ResponseData {
	return v.value
}

func (v *NullableExtractResources201ResponseData) Set(val *ExtractResources201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractResources201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractResources201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractResources201ResponseData(val *ExtractResources201ResponseData) *NullableExtractResources201ResponseData {
	return &NullableExtractResources201ResponseData{value: val, isSet: true}
}

func (v NullableExtractResources201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractResources201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


