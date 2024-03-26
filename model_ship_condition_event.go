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

// checks if the ShipConditionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipConditionEvent{}

// ShipConditionEvent An event that represents damage or wear to a ship's reactor, frame, or engine, reducing the condition of the ship.
type ShipConditionEvent struct {
	Symbol string `json:"symbol"`
	Component string `json:"component"`
	// The name of the event.
	Name string `json:"name"`
	// A description of the event.
	Description string `json:"description"`
}

type _ShipConditionEvent ShipConditionEvent

// NewShipConditionEvent instantiates a new ShipConditionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipConditionEvent(symbol string, component string, name string, description string) *ShipConditionEvent {
	this := ShipConditionEvent{}
	this.Symbol = symbol
	this.Component = component
	this.Name = name
	this.Description = description
	return &this
}

// NewShipConditionEventWithDefaults instantiates a new ShipConditionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipConditionEventWithDefaults() *ShipConditionEvent {
	this := ShipConditionEvent{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipConditionEvent) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipConditionEvent) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipConditionEvent) SetSymbol(v string) {
	o.Symbol = v
}

// GetComponent returns the Component field value
func (o *ShipConditionEvent) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ShipConditionEvent) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ShipConditionEvent) SetComponent(v string) {
	o.Component = v
}

// GetName returns the Name field value
func (o *ShipConditionEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipConditionEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipConditionEvent) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ShipConditionEvent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipConditionEvent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipConditionEvent) SetDescription(v string) {
	o.Description = v
}

func (o ShipConditionEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipConditionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["component"] = o.Component
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *ShipConditionEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"component",
		"name",
		"description",
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

	varShipConditionEvent := _ShipConditionEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipConditionEvent)

	if err != nil {
		return err
	}

	*o = ShipConditionEvent(varShipConditionEvent)

	return err
}

type NullableShipConditionEvent struct {
	value *ShipConditionEvent
	isSet bool
}

func (v NullableShipConditionEvent) Get() *ShipConditionEvent {
	return v.value
}

func (v *NullableShipConditionEvent) Set(val *ShipConditionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableShipConditionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableShipConditionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipConditionEvent(val *ShipConditionEvent) *NullableShipConditionEvent {
	return &NullableShipConditionEvent{value: val, isSet: true}
}

func (v NullableShipConditionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipConditionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

