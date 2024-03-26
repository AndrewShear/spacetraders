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

// TradeSymbol The good's symbol.
type TradeSymbol string

// List of TradeSymbol
const (
	TRADESYMBOL_PRECIOUS_STONES TradeSymbol = "PRECIOUS_STONES"
	TRADESYMBOL_QUARTZ_SAND TradeSymbol = "QUARTZ_SAND"
	TRADESYMBOL_SILICON_CRYSTALS TradeSymbol = "SILICON_CRYSTALS"
	TRADESYMBOL_AMMONIA_ICE TradeSymbol = "AMMONIA_ICE"
	TRADESYMBOL_LIQUID_HYDROGEN TradeSymbol = "LIQUID_HYDROGEN"
	TRADESYMBOL_LIQUID_NITROGEN TradeSymbol = "LIQUID_NITROGEN"
	TRADESYMBOL_ICE_WATER TradeSymbol = "ICE_WATER"
	TRADESYMBOL_EXOTIC_MATTER TradeSymbol = "EXOTIC_MATTER"
	TRADESYMBOL_ADVANCED_CIRCUITRY TradeSymbol = "ADVANCED_CIRCUITRY"
	TRADESYMBOL_GRAVITON_EMITTERS TradeSymbol = "GRAVITON_EMITTERS"
	TRADESYMBOL_IRON TradeSymbol = "IRON"
	TRADESYMBOL_IRON_ORE TradeSymbol = "IRON_ORE"
	TRADESYMBOL_COPPER TradeSymbol = "COPPER"
	TRADESYMBOL_COPPER_ORE TradeSymbol = "COPPER_ORE"
	TRADESYMBOL_ALUMINUM TradeSymbol = "ALUMINUM"
	TRADESYMBOL_ALUMINUM_ORE TradeSymbol = "ALUMINUM_ORE"
	TRADESYMBOL_SILVER TradeSymbol = "SILVER"
	TRADESYMBOL_SILVER_ORE TradeSymbol = "SILVER_ORE"
	TRADESYMBOL_GOLD TradeSymbol = "GOLD"
	TRADESYMBOL_GOLD_ORE TradeSymbol = "GOLD_ORE"
	TRADESYMBOL_PLATINUM TradeSymbol = "PLATINUM"
	TRADESYMBOL_PLATINUM_ORE TradeSymbol = "PLATINUM_ORE"
	TRADESYMBOL_DIAMONDS TradeSymbol = "DIAMONDS"
	TRADESYMBOL_URANITE TradeSymbol = "URANITE"
	TRADESYMBOL_URANITE_ORE TradeSymbol = "URANITE_ORE"
	TRADESYMBOL_MERITIUM TradeSymbol = "MERITIUM"
	TRADESYMBOL_MERITIUM_ORE TradeSymbol = "MERITIUM_ORE"
	TRADESYMBOL_HYDROCARBON TradeSymbol = "HYDROCARBON"
	TRADESYMBOL_ANTIMATTER TradeSymbol = "ANTIMATTER"
	TRADESYMBOL_FAB_MATS TradeSymbol = "FAB_MATS"
	TRADESYMBOL_FERTILIZERS TradeSymbol = "FERTILIZERS"
	TRADESYMBOL_FABRICS TradeSymbol = "FABRICS"
	TRADESYMBOL_FOOD TradeSymbol = "FOOD"
	TRADESYMBOL_JEWELRY TradeSymbol = "JEWELRY"
	TRADESYMBOL_MACHINERY TradeSymbol = "MACHINERY"
	TRADESYMBOL_FIREARMS TradeSymbol = "FIREARMS"
	TRADESYMBOL_ASSAULT_RIFLES TradeSymbol = "ASSAULT_RIFLES"
	TRADESYMBOL_MILITARY_EQUIPMENT TradeSymbol = "MILITARY_EQUIPMENT"
	TRADESYMBOL_EXPLOSIVES TradeSymbol = "EXPLOSIVES"
	TRADESYMBOL_LAB_INSTRUMENTS TradeSymbol = "LAB_INSTRUMENTS"
	TRADESYMBOL_AMMUNITION TradeSymbol = "AMMUNITION"
	TRADESYMBOL_ELECTRONICS TradeSymbol = "ELECTRONICS"
	TRADESYMBOL_SHIP_PLATING TradeSymbol = "SHIP_PLATING"
	TRADESYMBOL_SHIP_PARTS TradeSymbol = "SHIP_PARTS"
	TRADESYMBOL_EQUIPMENT TradeSymbol = "EQUIPMENT"
	TRADESYMBOL_FUEL TradeSymbol = "FUEL"
	TRADESYMBOL_MEDICINE TradeSymbol = "MEDICINE"
	TRADESYMBOL_DRUGS TradeSymbol = "DRUGS"
	TRADESYMBOL_CLOTHING TradeSymbol = "CLOTHING"
	TRADESYMBOL_MICROPROCESSORS TradeSymbol = "MICROPROCESSORS"
	TRADESYMBOL_PLASTICS TradeSymbol = "PLASTICS"
	TRADESYMBOL_POLYNUCLEOTIDES TradeSymbol = "POLYNUCLEOTIDES"
	TRADESYMBOL_BIOCOMPOSITES TradeSymbol = "BIOCOMPOSITES"
	TRADESYMBOL_QUANTUM_STABILIZERS TradeSymbol = "QUANTUM_STABILIZERS"
	TRADESYMBOL_NANOBOTS TradeSymbol = "NANOBOTS"
	TRADESYMBOL_AI_MAINFRAMES TradeSymbol = "AI_MAINFRAMES"
	TRADESYMBOL_QUANTUM_DRIVES TradeSymbol = "QUANTUM_DRIVES"
	TRADESYMBOL_ROBOTIC_DRONES TradeSymbol = "ROBOTIC_DRONES"
	TRADESYMBOL_CYBER_IMPLANTS TradeSymbol = "CYBER_IMPLANTS"
	TRADESYMBOL_GENE_THERAPEUTICS TradeSymbol = "GENE_THERAPEUTICS"
	TRADESYMBOL_NEURAL_CHIPS TradeSymbol = "NEURAL_CHIPS"
	TRADESYMBOL_MOOD_REGULATORS TradeSymbol = "MOOD_REGULATORS"
	TRADESYMBOL_VIRAL_AGENTS TradeSymbol = "VIRAL_AGENTS"
	TRADESYMBOL_MICRO_FUSION_GENERATORS TradeSymbol = "MICRO_FUSION_GENERATORS"
	TRADESYMBOL_SUPERGRAINS TradeSymbol = "SUPERGRAINS"
	TRADESYMBOL_LASER_RIFLES TradeSymbol = "LASER_RIFLES"
	TRADESYMBOL_HOLOGRAPHICS TradeSymbol = "HOLOGRAPHICS"
	TRADESYMBOL_SHIP_SALVAGE TradeSymbol = "SHIP_SALVAGE"
	TRADESYMBOL_RELIC_TECH TradeSymbol = "RELIC_TECH"
	TRADESYMBOL_NOVEL_LIFEFORMS TradeSymbol = "NOVEL_LIFEFORMS"
	TRADESYMBOL_BOTANICAL_SPECIMENS TradeSymbol = "BOTANICAL_SPECIMENS"
	TRADESYMBOL_CULTURAL_ARTIFACTS TradeSymbol = "CULTURAL_ARTIFACTS"
	TRADESYMBOL_FRAME_PROBE TradeSymbol = "FRAME_PROBE"
	TRADESYMBOL_FRAME_DRONE TradeSymbol = "FRAME_DRONE"
	TRADESYMBOL_FRAME_INTERCEPTOR TradeSymbol = "FRAME_INTERCEPTOR"
	TRADESYMBOL_FRAME_RACER TradeSymbol = "FRAME_RACER"
	TRADESYMBOL_FRAME_FIGHTER TradeSymbol = "FRAME_FIGHTER"
	TRADESYMBOL_FRAME_FRIGATE TradeSymbol = "FRAME_FRIGATE"
	TRADESYMBOL_FRAME_SHUTTLE TradeSymbol = "FRAME_SHUTTLE"
	TRADESYMBOL_FRAME_EXPLORER TradeSymbol = "FRAME_EXPLORER"
	TRADESYMBOL_FRAME_MINER TradeSymbol = "FRAME_MINER"
	TRADESYMBOL_FRAME_LIGHT_FREIGHTER TradeSymbol = "FRAME_LIGHT_FREIGHTER"
	TRADESYMBOL_FRAME_HEAVY_FREIGHTER TradeSymbol = "FRAME_HEAVY_FREIGHTER"
	TRADESYMBOL_FRAME_TRANSPORT TradeSymbol = "FRAME_TRANSPORT"
	TRADESYMBOL_FRAME_DESTROYER TradeSymbol = "FRAME_DESTROYER"
	TRADESYMBOL_FRAME_CRUISER TradeSymbol = "FRAME_CRUISER"
	TRADESYMBOL_FRAME_CARRIER TradeSymbol = "FRAME_CARRIER"
	TRADESYMBOL_REACTOR_SOLAR_I TradeSymbol = "REACTOR_SOLAR_I"
	TRADESYMBOL_REACTOR_FUSION_I TradeSymbol = "REACTOR_FUSION_I"
	TRADESYMBOL_REACTOR_FISSION_I TradeSymbol = "REACTOR_FISSION_I"
	TRADESYMBOL_REACTOR_CHEMICAL_I TradeSymbol = "REACTOR_CHEMICAL_I"
	TRADESYMBOL_REACTOR_ANTIMATTER_I TradeSymbol = "REACTOR_ANTIMATTER_I"
	TRADESYMBOL_ENGINE_IMPULSE_DRIVE_I TradeSymbol = "ENGINE_IMPULSE_DRIVE_I"
	TRADESYMBOL_ENGINE_ION_DRIVE_I TradeSymbol = "ENGINE_ION_DRIVE_I"
	TRADESYMBOL_ENGINE_ION_DRIVE_II TradeSymbol = "ENGINE_ION_DRIVE_II"
	TRADESYMBOL_ENGINE_HYPER_DRIVE_I TradeSymbol = "ENGINE_HYPER_DRIVE_I"
	TRADESYMBOL_MODULE_MINERAL_PROCESSOR_I TradeSymbol = "MODULE_MINERAL_PROCESSOR_I"
	TRADESYMBOL_MODULE_GAS_PROCESSOR_I TradeSymbol = "MODULE_GAS_PROCESSOR_I"
	TRADESYMBOL_MODULE_CARGO_HOLD_I TradeSymbol = "MODULE_CARGO_HOLD_I"
	TRADESYMBOL_MODULE_CARGO_HOLD_II TradeSymbol = "MODULE_CARGO_HOLD_II"
	TRADESYMBOL_MODULE_CARGO_HOLD_III TradeSymbol = "MODULE_CARGO_HOLD_III"
	TRADESYMBOL_MODULE_CREW_QUARTERS_I TradeSymbol = "MODULE_CREW_QUARTERS_I"
	TRADESYMBOL_MODULE_ENVOY_QUARTERS_I TradeSymbol = "MODULE_ENVOY_QUARTERS_I"
	TRADESYMBOL_MODULE_PASSENGER_CABIN_I TradeSymbol = "MODULE_PASSENGER_CABIN_I"
	TRADESYMBOL_MODULE_MICRO_REFINERY_I TradeSymbol = "MODULE_MICRO_REFINERY_I"
	TRADESYMBOL_MODULE_SCIENCE_LAB_I TradeSymbol = "MODULE_SCIENCE_LAB_I"
	TRADESYMBOL_MODULE_JUMP_DRIVE_I TradeSymbol = "MODULE_JUMP_DRIVE_I"
	TRADESYMBOL_MODULE_JUMP_DRIVE_II TradeSymbol = "MODULE_JUMP_DRIVE_II"
	TRADESYMBOL_MODULE_JUMP_DRIVE_III TradeSymbol = "MODULE_JUMP_DRIVE_III"
	TRADESYMBOL_MODULE_WARP_DRIVE_I TradeSymbol = "MODULE_WARP_DRIVE_I"
	TRADESYMBOL_MODULE_WARP_DRIVE_II TradeSymbol = "MODULE_WARP_DRIVE_II"
	TRADESYMBOL_MODULE_WARP_DRIVE_III TradeSymbol = "MODULE_WARP_DRIVE_III"
	TRADESYMBOL_MODULE_SHIELD_GENERATOR_I TradeSymbol = "MODULE_SHIELD_GENERATOR_I"
	TRADESYMBOL_MODULE_SHIELD_GENERATOR_II TradeSymbol = "MODULE_SHIELD_GENERATOR_II"
	TRADESYMBOL_MODULE_ORE_REFINERY_I TradeSymbol = "MODULE_ORE_REFINERY_I"
	TRADESYMBOL_MODULE_FUEL_REFINERY_I TradeSymbol = "MODULE_FUEL_REFINERY_I"
	TRADESYMBOL_MOUNT_GAS_SIPHON_I TradeSymbol = "MOUNT_GAS_SIPHON_I"
	TRADESYMBOL_MOUNT_GAS_SIPHON_II TradeSymbol = "MOUNT_GAS_SIPHON_II"
	TRADESYMBOL_MOUNT_GAS_SIPHON_III TradeSymbol = "MOUNT_GAS_SIPHON_III"
	TRADESYMBOL_MOUNT_SURVEYOR_I TradeSymbol = "MOUNT_SURVEYOR_I"
	TRADESYMBOL_MOUNT_SURVEYOR_II TradeSymbol = "MOUNT_SURVEYOR_II"
	TRADESYMBOL_MOUNT_SURVEYOR_III TradeSymbol = "MOUNT_SURVEYOR_III"
	TRADESYMBOL_MOUNT_SENSOR_ARRAY_I TradeSymbol = "MOUNT_SENSOR_ARRAY_I"
	TRADESYMBOL_MOUNT_SENSOR_ARRAY_II TradeSymbol = "MOUNT_SENSOR_ARRAY_II"
	TRADESYMBOL_MOUNT_SENSOR_ARRAY_III TradeSymbol = "MOUNT_SENSOR_ARRAY_III"
	TRADESYMBOL_MOUNT_MINING_LASER_I TradeSymbol = "MOUNT_MINING_LASER_I"
	TRADESYMBOL_MOUNT_MINING_LASER_II TradeSymbol = "MOUNT_MINING_LASER_II"
	TRADESYMBOL_MOUNT_MINING_LASER_III TradeSymbol = "MOUNT_MINING_LASER_III"
	TRADESYMBOL_MOUNT_LASER_CANNON_I TradeSymbol = "MOUNT_LASER_CANNON_I"
	TRADESYMBOL_MOUNT_MISSILE_LAUNCHER_I TradeSymbol = "MOUNT_MISSILE_LAUNCHER_I"
	TRADESYMBOL_MOUNT_TURRET_I TradeSymbol = "MOUNT_TURRET_I"
	TRADESYMBOL_SHIP_PROBE TradeSymbol = "SHIP_PROBE"
	TRADESYMBOL_SHIP_MINING_DRONE TradeSymbol = "SHIP_MINING_DRONE"
	TRADESYMBOL_SHIP_SIPHON_DRONE TradeSymbol = "SHIP_SIPHON_DRONE"
	TRADESYMBOL_SHIP_INTERCEPTOR TradeSymbol = "SHIP_INTERCEPTOR"
	TRADESYMBOL_SHIP_LIGHT_HAULER TradeSymbol = "SHIP_LIGHT_HAULER"
	TRADESYMBOL_SHIP_COMMAND_FRIGATE TradeSymbol = "SHIP_COMMAND_FRIGATE"
	TRADESYMBOL_SHIP_EXPLORER TradeSymbol = "SHIP_EXPLORER"
	TRADESYMBOL_SHIP_HEAVY_FREIGHTER TradeSymbol = "SHIP_HEAVY_FREIGHTER"
	TRADESYMBOL_SHIP_LIGHT_SHUTTLE TradeSymbol = "SHIP_LIGHT_SHUTTLE"
	TRADESYMBOL_SHIP_ORE_HOUND TradeSymbol = "SHIP_ORE_HOUND"
	TRADESYMBOL_SHIP_REFINING_FREIGHTER TradeSymbol = "SHIP_REFINING_FREIGHTER"
	TRADESYMBOL_SHIP_SURVEYOR TradeSymbol = "SHIP_SURVEYOR"
)

// All allowed values of TradeSymbol enum
var AllowedTradeSymbolEnumValues = []TradeSymbol{
	"PRECIOUS_STONES",
	"QUARTZ_SAND",
	"SILICON_CRYSTALS",
	"AMMONIA_ICE",
	"LIQUID_HYDROGEN",
	"LIQUID_NITROGEN",
	"ICE_WATER",
	"EXOTIC_MATTER",
	"ADVANCED_CIRCUITRY",
	"GRAVITON_EMITTERS",
	"IRON",
	"IRON_ORE",
	"COPPER",
	"COPPER_ORE",
	"ALUMINUM",
	"ALUMINUM_ORE",
	"SILVER",
	"SILVER_ORE",
	"GOLD",
	"GOLD_ORE",
	"PLATINUM",
	"PLATINUM_ORE",
	"DIAMONDS",
	"URANITE",
	"URANITE_ORE",
	"MERITIUM",
	"MERITIUM_ORE",
	"HYDROCARBON",
	"ANTIMATTER",
	"FAB_MATS",
	"FERTILIZERS",
	"FABRICS",
	"FOOD",
	"JEWELRY",
	"MACHINERY",
	"FIREARMS",
	"ASSAULT_RIFLES",
	"MILITARY_EQUIPMENT",
	"EXPLOSIVES",
	"LAB_INSTRUMENTS",
	"AMMUNITION",
	"ELECTRONICS",
	"SHIP_PLATING",
	"SHIP_PARTS",
	"EQUIPMENT",
	"FUEL",
	"MEDICINE",
	"DRUGS",
	"CLOTHING",
	"MICROPROCESSORS",
	"PLASTICS",
	"POLYNUCLEOTIDES",
	"BIOCOMPOSITES",
	"QUANTUM_STABILIZERS",
	"NANOBOTS",
	"AI_MAINFRAMES",
	"QUANTUM_DRIVES",
	"ROBOTIC_DRONES",
	"CYBER_IMPLANTS",
	"GENE_THERAPEUTICS",
	"NEURAL_CHIPS",
	"MOOD_REGULATORS",
	"VIRAL_AGENTS",
	"MICRO_FUSION_GENERATORS",
	"SUPERGRAINS",
	"LASER_RIFLES",
	"HOLOGRAPHICS",
	"SHIP_SALVAGE",
	"RELIC_TECH",
	"NOVEL_LIFEFORMS",
	"BOTANICAL_SPECIMENS",
	"CULTURAL_ARTIFACTS",
	"FRAME_PROBE",
	"FRAME_DRONE",
	"FRAME_INTERCEPTOR",
	"FRAME_RACER",
	"FRAME_FIGHTER",
	"FRAME_FRIGATE",
	"FRAME_SHUTTLE",
	"FRAME_EXPLORER",
	"FRAME_MINER",
	"FRAME_LIGHT_FREIGHTER",
	"FRAME_HEAVY_FREIGHTER",
	"FRAME_TRANSPORT",
	"FRAME_DESTROYER",
	"FRAME_CRUISER",
	"FRAME_CARRIER",
	"REACTOR_SOLAR_I",
	"REACTOR_FUSION_I",
	"REACTOR_FISSION_I",
	"REACTOR_CHEMICAL_I",
	"REACTOR_ANTIMATTER_I",
	"ENGINE_IMPULSE_DRIVE_I",
	"ENGINE_ION_DRIVE_I",
	"ENGINE_ION_DRIVE_II",
	"ENGINE_HYPER_DRIVE_I",
	"MODULE_MINERAL_PROCESSOR_I",
	"MODULE_GAS_PROCESSOR_I",
	"MODULE_CARGO_HOLD_I",
	"MODULE_CARGO_HOLD_II",
	"MODULE_CARGO_HOLD_III",
	"MODULE_CREW_QUARTERS_I",
	"MODULE_ENVOY_QUARTERS_I",
	"MODULE_PASSENGER_CABIN_I",
	"MODULE_MICRO_REFINERY_I",
	"MODULE_SCIENCE_LAB_I",
	"MODULE_JUMP_DRIVE_I",
	"MODULE_JUMP_DRIVE_II",
	"MODULE_JUMP_DRIVE_III",
	"MODULE_WARP_DRIVE_I",
	"MODULE_WARP_DRIVE_II",
	"MODULE_WARP_DRIVE_III",
	"MODULE_SHIELD_GENERATOR_I",
	"MODULE_SHIELD_GENERATOR_II",
	"MODULE_ORE_REFINERY_I",
	"MODULE_FUEL_REFINERY_I",
	"MOUNT_GAS_SIPHON_I",
	"MOUNT_GAS_SIPHON_II",
	"MOUNT_GAS_SIPHON_III",
	"MOUNT_SURVEYOR_I",
	"MOUNT_SURVEYOR_II",
	"MOUNT_SURVEYOR_III",
	"MOUNT_SENSOR_ARRAY_I",
	"MOUNT_SENSOR_ARRAY_II",
	"MOUNT_SENSOR_ARRAY_III",
	"MOUNT_MINING_LASER_I",
	"MOUNT_MINING_LASER_II",
	"MOUNT_MINING_LASER_III",
	"MOUNT_LASER_CANNON_I",
	"MOUNT_MISSILE_LAUNCHER_I",
	"MOUNT_TURRET_I",
	"SHIP_PROBE",
	"SHIP_MINING_DRONE",
	"SHIP_SIPHON_DRONE",
	"SHIP_INTERCEPTOR",
	"SHIP_LIGHT_HAULER",
	"SHIP_COMMAND_FRIGATE",
	"SHIP_EXPLORER",
	"SHIP_HEAVY_FREIGHTER",
	"SHIP_LIGHT_SHUTTLE",
	"SHIP_ORE_HOUND",
	"SHIP_REFINING_FREIGHTER",
	"SHIP_SURVEYOR",
}

func (v *TradeSymbol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TradeSymbol(value)
	for _, existing := range AllowedTradeSymbolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TradeSymbol", value)
}

// NewTradeSymbolFromValue returns a pointer to a valid TradeSymbol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradeSymbolFromValue(v string) (*TradeSymbol, error) {
	ev := TradeSymbol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TradeSymbol: valid values are %v", v, AllowedTradeSymbolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TradeSymbol) IsValid() bool {
	for _, existing := range AllowedTradeSymbolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TradeSymbol value
func (v TradeSymbol) Ptr() *TradeSymbol {
	return &v
}

type NullableTradeSymbol struct {
	value *TradeSymbol
	isSet bool
}

func (v NullableTradeSymbol) Get() *TradeSymbol {
	return v.value
}

func (v *NullableTradeSymbol) Set(val *TradeSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeSymbol(val *TradeSymbol) *NullableTradeSymbol {
	return &NullableTradeSymbol{value: val, isSet: true}
}

func (v NullableTradeSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

