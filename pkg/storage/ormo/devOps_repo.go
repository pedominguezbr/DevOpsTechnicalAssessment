package ormo

import (
	"database/sql"
	"framework-go/pkg/features/devOps"
)

type devOpsRepo struct {
	db      []*sql.DB
	connSix *sql.DB
}

// NewdevOpsRepo funcion instancia
func NewdevOpsRepo(db []*sql.DB) devOps.Repository {
	return &devOpsRepo{db: db, connSix: db[0]}
}

func (s *devOpsRepo) DevOps(data *devOps.RequestDevops) (err error) {

	return nil
}

// NewNullInt64 return null value
func NewNullInt64(val *int64) sql.NullInt64 {
	if val != nil {
		return sql.NullInt64{Int64: *val, Valid: true}
	}

	return sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
}
