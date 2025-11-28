package geofencing_zones

import (
	geojson "github.com/paulmach/orb/geojson"
	"github.com/phd-kerger/gbfs-go-adapter/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v3.0/header"
)

// Describes geofencing zones and their associated rules and attributes (added in v2.1-RC).
type GeofencingZones struct {
	header.HeaderStruct
	// Array that contains geofencing information for the system.
	Data Data `json:"data"`
}

// Array that contains geofencing information for the system.
type Data struct {
	// Each geofenced zone and its associated rules and attributes is described as an object
	// within the array of features.
	GeofencingZones *geojson.FeatureCollection `json:"geofencing_zones"`
	// Array of Rule objects defining restrictions that apply globally in all areas as the
	// default restrictions, except where overridden with an explicit geofencing zone.
	GlobalRules []GlobalRule `json:"global_rules"`
}

type GlobalRule struct {
	// What is the maximum speed allowed, in kilometers per hour?
	MaximumSpeedKph *uint64 `json:"maximum_speed_kph,omitempty"`
	// Is the ride allowed to end in this zone?
	RideEndAllowed bool `json:"ride_end_allowed"`
	// Is the ride allowed to start in this zone?
	RideStartAllowed bool `json:"ride_start_allowed"`
	// Is the ride allowed to travel through this zone?
	RideThroughAllowed bool `json:"ride_through_allowed"`
	// Vehicle MUST be parked at stations defined in station_information.json within this
	// geofence zone
	StationParking *bool `json:"station_parking,omitempty"`
	// Array of vehicle type IDs for which these restrictions apply.
	VehicleTypeIDS []common.ID `json:"vehicle_type_ids,omitempty"`
}
