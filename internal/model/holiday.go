package model

type Holiday struct {
	Date   Date
	Name   string
	Type   HolidayReason
	Reason string
}

type HolidayReason int

const (
	HolidayReasonWeekHoliday HolidayReason = iota
	HolidayReasonFederalHoliday
	HolidayReasonStateHoliday
)
