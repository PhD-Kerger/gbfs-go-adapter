package station_information

import (
	station_information_v30 "github.com/phd-kerger/gbfs-go-adapter/models/v3.0/station_information"
)

// List of all stations, their capacities and locations. REQUIRED of systems utilizing docks.
type StationInformation struct {
	station_information_v30.StationInformation
	// Array that contains one object per station as defined below.
	Data Data `json:"data"`
}

type Data struct {
	Stations []Station `json:"stations"`
}

type Station struct {
	station_information_v30.Station
	City *string `json:"city,omitempty"`
}
