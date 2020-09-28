package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpHandler "github.com/jonathangunawan/test-bobobox/handler/http"
	"github.com/jonathangunawan/test-bobobox/infra/config"
	"github.com/jonathangunawan/test-bobobox/infra/database"
	"github.com/jonathangunawan/test-bobobox/repository"
	"github.com/jonathangunawan/test-bobobox/usecase"
)

var cfgFilePath string

func init() {
	cfgFilePath = "files/config/config.json"
}

func main() {
	r := setupHTTPServer()
	log.Println("Serving at port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func setupHTTPServer() *mux.Router {
	//Read Config
	icfg := config.New(cfgFilePath)
	cfg, err := icfg.ReadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to read Config ERR: %s", err.Error()))
	}

	//Init Database
	dbPG, err := database.NewPostgreDB(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to init Database ERR: %s", err.Error()))
	}

	//Init Repository
	repo := repository.New(dbPG)

	//Init UseCase
	uc := usecase.New(repo)

	//Init HTTP Handler
	h := httpHandler.New(uc)

	//Register Router
	r := mux.NewRouter()

	//Comments
	r.HandleFunc("/booking", h.BookingHandler).
		Methods(http.MethodPost)

	return r
}
