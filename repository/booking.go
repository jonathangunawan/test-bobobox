package repository

import (
	"github.com/jonathangunawan/test-bobobox/entity"
)

func (r RepoImpl) InsertBooking(input entity.BookingDB) error {
	_, err := r.db.Conn.NamedExec(queryInsertBooking, input)
	if err != nil {
		return err
	}

	return nil
}
