package vehicle_status

import (
	vehicle_status_v30 "github.com/phd-kerger/gbfs-go-adapter/models/v3.0/vehicle_status"
)

// Describes the vehicles that are available for rent (as of v3.0, formerly
// free_bike_status).
type VehicleStatus struct {
	vehicle_status_v30.VehicleStatus
}
