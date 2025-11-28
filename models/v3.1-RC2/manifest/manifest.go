package manifest

import (
	"github.com/PhD-Kerger/gbfs-go-adapter/models/common"
	manifest_v30 "github.com/PhD-Kerger/gbfs-go-adapter/models/v3.0/manifest"
	"github.com/paulmach/orb/geojson"
)

// An index of gbfs.json URLs for each GBFS data set produced by a publisher. A single
// instance of this file should be published at a single stable URL, for example:
// https://example.com/gbfs/manifest.json.
type Manifest struct {
	manifest_v30.Manifest
	Data Data `json:"data"`
}

type Data struct {
	// An array of objects, each containing the keys below.
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	manifest_v30.Dataset
	// A GeoJSON MultiPolygon that describes the operating area.
	// If area is supplied, then the record describes the general operating area of the system for the purpose of discovery.
	// Geographic details of the system's operating restrictions must be explicitly specified using station locations and geofencing zones, where appropriate.
	Area *geojson.MultiPolygon `json:"area,omitempty"`
	// The ISO 3166-1 alpha-2 country code of the operating area. The field MUST NOT be specified if the operating area spans multiple countries.
	CountryCode *common.CountryCode `json:"country_code,omitempty"`
}
