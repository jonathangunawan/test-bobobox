package usecase

import (
	"github.com/jonathangunawan/test-bobobox/entity"
	repo "github.com/jonathangunawan/test-bobobox/repository"
)

type UseCaseItf interface {
	CreateResi(entity.BookingHTTP) (string, error)
}

type UseCaseImpl struct {
	repo repo.RepoItf
}

func New(repo repo.RepoItf) UseCaseItf {
	return UseCaseImpl{
		repo: repo,
	}
}
