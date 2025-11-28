package station_status

import (
	station_status_v30 "github.com/PhD-Kerger/gbfs-go-adapter/models/v3.0/station_status"
)

// Describes the capacity and rental availability of the station
type StationStatus struct {
	station_status_v30.StationStatus
}
