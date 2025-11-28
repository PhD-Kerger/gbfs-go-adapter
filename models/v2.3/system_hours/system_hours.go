package systemhours

import (
	"github.com/phd-kerger/gbfs-go-adapter/models/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v2.3/header"
)

// Describes the system hours of operation.
type SystemHours struct {
	header.HeaderStruct
	// Array that contains system hours of operations.
	Data Data `json:"data"`
}

type Data struct {
	RentalHours []RentalHour `json:"rental_hours"`
}

type RentalHour struct {
	// Array of member and nonmember value(s) indicating that this set of rental hours applies to either members or non-members only.
	UserTypes []common.UserType `json:"user_types"`
	// An array of abbreviations (first 3 letters) of English names of the days of the week for which this object applies.
	Days []common.Day `json:"days"`
	// Start time for the hours of operation of the system.
	StartTime string `json:"start_time"`
	// End time for the hours of operation of the system.
	EndTime string `json:"end_time"`
}
