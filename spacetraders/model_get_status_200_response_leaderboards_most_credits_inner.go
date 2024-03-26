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

// checks if the GetStatus200ResponseLeaderboardsMostCreditsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200ResponseLeaderboardsMostCreditsInner{}

// GetStatus200ResponseLeaderboardsMostCreditsInner struct for GetStatus200ResponseLeaderboardsMostCreditsInner
type GetStatus200ResponseLeaderboardsMostCreditsInner struct {
	// Symbol of the agent.
	AgentSymbol string `json:"agentSymbol"`
	// Amount of credits.
	Credits int64 `json:"credits"`
}

type _GetStatus200ResponseLeaderboardsMostCreditsInner GetStatus200ResponseLeaderboardsMostCreditsInner

// NewGetStatus200ResponseLeaderboardsMostCreditsInner instantiates a new GetStatus200ResponseLeaderboardsMostCreditsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200ResponseLeaderboardsMostCreditsInner(agentSymbol string, credits int64) *GetStatus200ResponseLeaderboardsMostCreditsInner {
	this := GetStatus200ResponseLeaderboardsMostCreditsInner{}
	this.AgentSymbol = agentSymbol
	this.Credits = credits
	return &this
}

// NewGetStatus200ResponseLeaderboardsMostCreditsInnerWithDefaults instantiates a new GetStatus200ResponseLeaderboardsMostCreditsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseLeaderboardsMostCreditsInnerWithDefaults() *GetStatus200ResponseLeaderboardsMostCreditsInner {
	this := GetStatus200ResponseLeaderboardsMostCreditsInner{}
	return &this
}

// GetAgentSymbol returns the AgentSymbol field value
func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) GetAgentSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentSymbol
}

// GetAgentSymbolOk returns a tuple with the AgentSymbol field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) GetAgentSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentSymbol, true
}

// SetAgentSymbol sets field value
func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) SetAgentSymbol(v string) {
	o.AgentSymbol = v
}

// GetCredits returns the Credits field value
func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) GetCredits() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) GetCreditsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credits, true
}

// SetCredits sets field value
func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) SetCredits(v int64) {
	o.Credits = v
}

func (o GetStatus200ResponseLeaderboardsMostCreditsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200ResponseLeaderboardsMostCreditsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentSymbol"] = o.AgentSymbol
	toSerialize["credits"] = o.Credits
	return toSerialize, nil
}

func (o *GetStatus200ResponseLeaderboardsMostCreditsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agentSymbol",
		"credits",
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

	varGetStatus200ResponseLeaderboardsMostCreditsInner := _GetStatus200ResponseLeaderboardsMostCreditsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatus200ResponseLeaderboardsMostCreditsInner)

	if err != nil {
		return err
	}

	*o = GetStatus200ResponseLeaderboardsMostCreditsInner(varGetStatus200ResponseLeaderboardsMostCreditsInner)

	return err
}

type NullableGetStatus200ResponseLeaderboardsMostCreditsInner struct {
	value *GetStatus200ResponseLeaderboardsMostCreditsInner
	isSet bool
}

func (v NullableGetStatus200ResponseLeaderboardsMostCreditsInner) Get() *GetStatus200ResponseLeaderboardsMostCreditsInner {
	return v.value
}

func (v *NullableGetStatus200ResponseLeaderboardsMostCreditsInner) Set(val *GetStatus200ResponseLeaderboardsMostCreditsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200ResponseLeaderboardsMostCreditsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200ResponseLeaderboardsMostCreditsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200ResponseLeaderboardsMostCreditsInner(val *GetStatus200ResponseLeaderboardsMostCreditsInner) *NullableGetStatus200ResponseLeaderboardsMostCreditsInner {
	return &NullableGetStatus200ResponseLeaderboardsMostCreditsInner{value: val, isSet: true}
}

func (v NullableGetStatus200ResponseLeaderboardsMostCreditsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200ResponseLeaderboardsMostCreditsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


