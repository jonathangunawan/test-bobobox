package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/jonathangunawan/test-bobobox/entity"
	_ "github.com/lib/pq"
)

type PGImpl struct {
	Conn *sqlx.DB
}

func NewPostgreDB(cfg entity.Config) (PGImpl, error) {
	db, err := sqlx.Open(cfg.DBType, cfg.DBConString)
	if err != nil {
		return PGImpl{}, err
	}

	return PGImpl{
		Conn: db,
	}, nil
}
