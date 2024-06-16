package bridging_day_suggester

import (
	"context"
	"fmt"

	"github.com/yegorkir/bridgingday/internal/model"
	"github.com/yegorkir/bridgingday/internal/services/bridging_day_suggester/dto"
	"github.com/yegorkir/bridgingday/pkg/log"
)

//go:generate mockery --name holidayRepository --exported --with-expecter
type holidayRepository interface {
	GetHolidays(ctx context.Context, lid model.LocationID, period model.DatePeriod) ([]model.Holiday, error)
}

type Service struct {
	// yearRepository yearRepository
	holidayRepository holidayRepository

	log *log.Logger
}

func NewService(holidayRepository holidayRepository, log *log.Logger) *Service {
	return &Service{
		holidayRepository: holidayRepository,
		log:               log,
	}
}

func (s *Service) Suggest(in *dto.SuggestIn) (*dto.SuggestOut, error) {
	// calendar, err := s.yearRepository.CalendarByLocation(in.LID, year)
	// if err != nil {
	// 	return nil, fmt.Errorf("bridging_day_suggester suggest: %w", err)
	// }

	return s.suggest(in)
}

func (s *Service) suggest(in *dto.SuggestIn) (*dto.SuggestOut, error) {
	holidays, err := s.holidayRepository.GetHolidays(context.Background(), in.LID, in.SearchingPeriod)
	if err != nil {
		return nil, fmt.Errorf("bridging_day_suggester suggest: %w", err)
	}

	if holidays[len(holidays) - 1].Date.More(in.SearchingPeriod.To) {
		s.log.Error("searching period is out of range", "period", in.SearchingPeriod, "holidays", holidays)
	}

	bridgingPeriods := findBridgingPeriods(in.SearchingPeriod, holidays)

	return &dto.SuggestOut{
		BridgingPeriods: bridgingPeriods,
		Holidays:        holidays,
	}, nil
}

// findBridgingPeriods searches periods between holidays
func findBridgingPeriods(searchingPeriod model.DatePeriod, holidays []model.Holiday) []model.DatePeriod {
	bridgingPeriods := make([]model.DatePeriod, 0)

	var from *model.Date
	if searchingPeriod.From != holidays[0].Date {
		from = &searchingPeriod.From
	}

	hIdx := 0
	for  {
 		if hIdx >= len(holidays) {
			break
		}
		if

	}

	return bridgingPeriods
}
