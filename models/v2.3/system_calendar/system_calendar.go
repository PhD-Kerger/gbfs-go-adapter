package systemcalendar

import "github.com/phd-kerger/gbfs-go-adapter/models/v2.3/header"

// Describes the operating calendar for a system.
type SystemCalendar struct {
	header.HeaderStruct
	// Array that contains operations calendar for the system.
	Data Data `json:"data"`
}

type Data struct {
	Calendars []Calendar `json:"calendars"`
}

type Calendar struct {
	// Starting month for the system operations.
	StartMonth uint64 `json:"start_month"`
	// Starting day for the system operations.
	StartDay uint64 `json:"start_day"`
	// Starting year for the system operations.
	StartYear *uint64 `json:"start_year,omitempty"`
	// End month for the system operations.
	EndMonth uint64 `json:"end_month"`
	// End day for the system operations.
	EndDay uint64 `json:"end_day"`
	// End year for the system operations.
	EndYear *uint64 `json:"end_year,omitempty"`
}
