package api

import (
	"log"
	"net/http"

	"github.com/kasariks/a_nameless_project_yet/sevice/tense"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	tenseHandler := tense.NewHandler()
	tenseHandler.RegisterRoutes(router)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
