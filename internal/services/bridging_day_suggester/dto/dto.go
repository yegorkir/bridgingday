package dto

import (
	"github.com/yegorkir/bridgingday/internal/model"
)

type SuggestIn struct {
	SearchingPeriod model.DatePeriod

	LID model.LocationID

	VacationDays int
}

type SuggestOut struct {
	BridgingPeriods []model.DatePeriod
	Holidays        []model.Holiday
}
