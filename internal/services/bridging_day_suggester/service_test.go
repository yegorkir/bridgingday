package bridging_day_suggester

import (
	"reflect"
	"testing"

	"github.com/yegorkir/bridgingday/internal/model"
)

func Test_findBridgingPeriods(t *testing.T) {
	type args struct {
		holidays []model.Holiday
	}
	tests := []struct {
		name     string
		holidays []model.Holiday
		want     []model.DatePeriod
	}{
		{
			name: "normal",
			holidays: []model.Holiday{
				{Date: model.NewDate(2020, 1, 1)},
				{Date: model.NewDate(2020, 1, 2)},
				{Date: model.NewDate(2020, 1, 3)},
				{Date: model.NewDate(2020, 1, 4)},
				{Date: model.NewDate(2020, 1, 5)},
				{Date: model.NewDate(2020, 1, 6)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBridgingPeriods(tt.holidays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBridgingPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
