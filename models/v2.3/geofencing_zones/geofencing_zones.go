package geofencing_zones

import (
	"github.com/PhD-Kerger/gbfs-go-adapter/models/v2.3/header"
	"github.com/paulmach/orb/geojson"
)

// Describes geofencing zones and their associated rules and attributes (added in v2.1-RC).
type GeofencingZones struct {
	header.HeaderStruct
	// Array that contains geofencing information for the system.
	Data Data `json:"data"`
}

// Array that contains geofencing information for the system.
type Data struct {
	// Each geofenced zone and its associated rules and attributes is described as an object  within the array of features.
	GeofencingZones *geojson.FeatureCollection `json:"geofencing_zones"`
}
