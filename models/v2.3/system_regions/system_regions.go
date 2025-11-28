package system_regions

import (
	"github.com/phd-kerger/gbfs-go-adapter/models/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v2.3/header"
)

// Describes regions for a system that is broken up by geographic or political region.
type SystemRegions struct {
	header.HeaderStruct
	// Describe regions for a system that is broken up by geographic or political region.
	Data Data `json:"data"`
}

// Describe regions for a system that is broken up by geographic or political region.
type Data struct {
	// Array of regions.
	Regions []Region `json:"regions"`
}

type Region struct {
	// Public name for this region.
	Name string `json:"name"`
	// identifier of the region.
	RegionID common.ID `json:"region_id"`
}
