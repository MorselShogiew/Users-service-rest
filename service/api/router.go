package api

import (
	"log"

	"github.com/gorilla/mux"

	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/repos"

	hv1 "MorselShogiew/Users-service-rest/service/api/handlers/v1"

	"MorselShogiew/Users-service-rest/service/usecases"
)

type Service struct {
	v1 *hv1.Handlers
	u  *usecases.Service
}

func New(l logger.Logger, r *repos.Repositories) *Service {
	u := usecases.New(r)
	return &Service{
		v1: hv1.New(u, l),
		u:  u,
	}
}

func (s *Service) Router(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	v1Auth := v1.PathPrefix("").Subrouter()

	v1Auth.HandleFunc("/add", s.v1.AddUser).Methods("POST")
	v1Auth.HandleFunc("/delete", s.v1.DeleteUser).Methods("POST")
	v1Auth.HandleFunc("/list", s.v1.GetUsers).Methods("GET")
}

func (s *Service) Start() error {
	log.Println(s.Name() + " started")

	return nil
}

func (s *Service) Stop() error {
	log.Println(s.Name() + " stopped")

	return nil
}

func (s *Service) Name() string {
	return "Resize Photo service"
}
