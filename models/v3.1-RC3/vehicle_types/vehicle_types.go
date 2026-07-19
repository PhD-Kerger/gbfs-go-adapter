package vehicle_types

import (
	vehicle_types_v30 "github.com/phd-kerger/gbfs-go-adapter/models/v3.0/vehicle_types"
)

// Describes the types of vehicles that System operator has available for rent (added in
// v2.1-RC).
type VehicleTypes struct {
	vehicle_types_v30.VehicleTypes
	Data Data `json:"data"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// Array that contains one object per vehicle type in the system as defined below.
	VehicleTypes []VehicleType `json:"vehicle_types"`
}

type VehicleType struct {
	vehicle_types_v30.VehicleType
	// Minimum age required to use this vehicle. Added in v3.1-RC3.
	MinimumAge *uint8 `json:"minimum_age,omitempty"`
}