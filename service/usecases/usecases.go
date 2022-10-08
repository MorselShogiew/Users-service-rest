package usecases

import (
	"MorselShogiew/Users-service-rest/repos"
)

type Service struct {
	APIRepo repos.DBRepo
}

func New(r *repos.Repositories) *Service {
	return &Service{
		r.DBRepo,
	}
}
