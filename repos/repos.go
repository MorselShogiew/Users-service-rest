package repos

import (
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/provider"
)

type Repositories struct {
	DBRepo DBRepo
	Cache  Cache
}

func New(p provider.Provider, l logger.Logger) *Repositories {
	DBRepo := NewDBRepo(p, l)
	cache := NewRedisCache(p, l)
	return &Repositories{
		DBRepo,
		cache,
	}
}
