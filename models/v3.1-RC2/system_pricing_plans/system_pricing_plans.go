package system_pricing_plans

import (
	system_pricing_plans_v30 "github.com/phd-kerger/gbfs-go-adapter/models/v3.0/system_pricing_plans"
)

// Describes the pricing schemes of the system.
type SystemPricingPlans struct {
	system_pricing_plans_v30.SystemPricingPlans
	// Array that contains one object per plan as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per plan as defined below.
type Data struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	system_pricing_plans_v30.Plan
	//The cost, described as per minute rate, to reserve the vehicle prior to beginning a rental.
	// This amount is charged for each minute of the vehicle reservation until the rental is initiated,
	// or until the number of minutes defined in vehicle_types.json#default_reserve_time elapses,
	// whichever comes first. When using this field, you MUST declare a value in vehicle_types.json#default_reserve_time.
	// This field MUST NOT be combined in a single pricing plan with reservation_price_flat_rate.
	Reservation_price_per_min *float64 `json:"reservation_price_per_min,omitempty"`
	// The cost, described as a flat rate, to reserve the vehicle prior to beginning a rental.
	// This amount is charged once to reserve the vehicle for the duration of the time defined by
	// vehicle_types.json#default_reserve_time. When using this field, you MUST declare a value in
	// vehicle_types.json#default_reserve_time. This field MUST NOT be combined in a single pricing
	// plan with reservation_price_per_min.
	Reservation_price_flat_rate *float64 `json:"reservation_price_flat_rate,omitempty"`
	// Object defining a capped fare once a price threshold has been spent within a timeframe.
	// The same fare cap applies to each subsequent timeframe. For example, a fare capped at 15.00 CAD per 12-hour period.
	Fare_capping *FareCapping `json:"fare_capping,omitempty"`
}

type FareCapping struct {
	// Amount of time in minutes during which the fare is capped.
	Duration uint64 `json:"duration"`
	// The maximum fare threshold for the current timeframe, in the unit specified by currency.
	Price float64 `json:"price"`
}
