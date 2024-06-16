package queryer

import "github.com/jmoiron/sqlx"

type Queryer interface {
	sqlx.ExtContext
}
