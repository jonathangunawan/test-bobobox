package usecase

import (
	entity "github.com/jonathangunawan/test-bobobox/entity"
)

func (uc UseCaseImpl) CreateResi(input entity.BookingHTTP) (string, error) {
	resi := "RESI0001"
	price := int64(10000)

	data := entity.BookingDB{
		Resi:        resi,
		Price:       price,
		SenderName:  input.SenderName,
		SenderPhone: input.SenderPhone,
		Origin:      input.Origin,
		Destination: input.Destination,
		Weight:      input.Weight,
		Service:     input.Service,
	}

	err := uc.repo.InsertBooking(data)
	if err != nil {
		return "", err
	}

	return resi, nil
}
