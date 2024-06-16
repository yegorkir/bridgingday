package holidays

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/yegorkir/bridgingday/internal/model"
	"github.com/yegorkir/bridgingday/pkg/queryer"
)

type Repository struct {
	queryer queryer.Queryer
}

func New(queryer queryer.Queryer) *Repository {
	return &Repository{
		queryer: queryer,
	}
}

func (r *Repository) GetHolidays(ctx context.Context, lid model.LocationID, period model.DatePeriod) ([]model.Holiday, error) {
	query := `SELECT date, name, type, reason FROM holidays WHERE lid = $1 AND date >= $2 AND date <= $3`

	holidays := make([]model.Holiday, 0)
	err := sqlx.SelectContext(ctx, r.queryer, &holidays, query, lid, period.From, period.To)
	if err != nil {
		return nil, err
	}

	return holidays, nil
}
