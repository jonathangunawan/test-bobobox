package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jonathangunawan/test-bobobox/entity"
)

func (h HttpHandlerImpl) BookingHandler(w http.ResponseWriter, r *http.Request) {
	// Get Data from Request Body
	var req entity.BookingHTTP
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// JSON Unmarshal
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resi, err := h.uc.CreateResi(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := struct {
		Resi string `json:"resi"`
	}{
		Resi: resi,
	}

	res, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
