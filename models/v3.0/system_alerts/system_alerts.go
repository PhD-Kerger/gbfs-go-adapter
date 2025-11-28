package system_alerts

import (
	"github.com/phd-kerger/gbfs-go-adapter/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v3.0/header"
)

// Describes ad-hoc changes to the system.
type SystemAlerts struct {
	header.HeaderStruct

	// Array that contains ad-hoc alerts for the system.
	Data Data `json:"data"`
}

// Array that contains ad-hoc alerts for the system.
type Data struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	// Identifier for this alert.
	AlertID common.ID `json:"alert_id"`
	// Detailed description of the alert.
	Description []common.LocalizedString `json:"description,omitempty"`
	// Indicates the last time the info for the alert was updated.
	LastUpdated *common.Timestamp `json:"last_updated,omitempty"`
	// Array of identifiers of the regions for which this alert applies.
	RegionIDS []common.ID `json:"region_ids,omitempty"`
	// Array of identifiers of the stations for which this alert applies.
	StationIDS []common.ID `json:"station_ids,omitempty"`
	// A short summary of this alert to be displayed to the customer.
	Summary []common.LocalizedString `json:"summary"`
	// Array of objects indicating when the alert is in effect.
	Times []Time `json:"times,omitempty"`
	// Type of alert.
	Type common.Type `json:"type"`
	// URL where the customer can learn more information about this alert.
	URL []common.LocalizedURL `json:"url,omitempty"`
}

type Time struct {
	// End time of the alert.
	End *common.Timestamp `json:"end,omitempty"`
	// Start time of the alert.
	Start *common.Timestamp `json:"start,omitempty"`
}
