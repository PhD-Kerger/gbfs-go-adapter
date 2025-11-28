package system_information

import (
	system_information_v30 "github.com/PhD-Kerger/gbfs-go-adapter/models/v3.0/system_information"
)

// Details including system operator, system location, year implemented, URL, contact info,
// time zone.
type SystemInformation struct {
	system_information_v30.SystemInformation
}
