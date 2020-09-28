package http

import (
	"net/http"

	"github.com/jonathangunawan/test-bobobox/usecase"
)

type HttpHandlerItf interface {
	BookingHandler(w http.ResponseWriter, r *http.Request)
}

type HttpHandlerImpl struct {
	uc usecase.UseCaseItf
}

func New(uc usecase.UseCaseItf) HttpHandlerItf {
	return HttpHandlerImpl{
		uc: uc,
	}
}
