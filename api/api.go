package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router
}

const PathPrefix = "/v1"

func (a *API) RegisterRoutes() {
	fmt.Println("Creating the router...")
	a.Router = mux.NewRouter()
	s := a.Router.PathPrefix(PathPrefix).Subrouter()

	s.HandleFunc("/", HttpHandler(helloWorld)).Methods(http.MethodGet)
}

func respond(w http.ResponseWriter, payload interface{}) error {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
		return BadRequest{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func helloWorld(w http.ResponseWriter, r *http.Request) error {
	log.Println("Hello World!")
	return respond(w, "Hello World!")
}
