package api

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/MorselShogiew/Users-service-rest/logger"
	"github.com/MorselShogiew/Users-service-rest/repos"
	hv1 "github.com/MorselShogiew/Users-service-rest/service/api/handlers/v1"
	"github.com/MorselShogiew/Users-service-rest/service/usecases"
)

type ResizePhotoService struct {
	v1 *hv1.Handlers
	u  *usecases.ResizeService
}

func New(l logger.Logger, r *repos.Repositories) *ResizePhotoService {
	u := usecases.New(r)
	return &ResizePhotoService{
		v1: hv1.New(u, l),
		u:  u,
	}
}

func (s *ResizePhotoService) Router(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	v1Auth := v1.PathPrefix("").Subrouter()

	v1Auth.HandleFunc("/resize", s.v1.GetResizePhoto).Methods("GET")
	v1Auth.HandleFunc("/url", s.v1.PostUrl).Methods("POST")
}

func (s *ResizePhotoService) Start() error {
	log.Println(s.Name() + " started")

	return nil
}

func (s *ResizePhotoService) Stop() error {
	log.Println(s.Name() + " stopped")

	return nil
}

func (s *ResizePhotoService) Name() string {
	return "Resize Photo service"
}
