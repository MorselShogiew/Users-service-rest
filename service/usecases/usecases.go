package usecases

import (
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/provider"
	"MorselShogiew/Users-service-rest/repos"
)

type Service struct {
	APIRepo repos.DBRepo
	cache   repos.Cache
}

func New(p provider.Provider, l logger.Logger) *Service {
	r := repos.New(p, l)
	return &Service{
		r.DBRepo,
		r.Cache,
	}
}
