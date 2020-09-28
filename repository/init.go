package repository

import (
	"github.com/jonathangunawan/test-bobobox/entity"
	"github.com/jonathangunawan/test-bobobox/infra/database"
)

type RepoItf interface {
	InsertBooking(entity.BookingDB) error
}

type RepoImpl struct {
	db database.PGImpl
}

func New(db database.PGImpl) RepoItf {
	return RepoImpl{
		db: db,
	}
}
