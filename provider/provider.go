package provider

import (
	"net/http"
	"time"

	"MorselShogiew/Users-service-rest/config"
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/provider/database"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Provider interface {
	GetDBConn() *sqlx.DB
	GetCacheClient() *redis.Client
	Close()
}

type provider struct {
	resources *resources
	apis      *apis
	c         *http.Client
	l         logger.Logger
}

// resources that should be closed manually at the end
type resources struct {
	db          *sqlx.DB
	cacheClient *redis.Client
}

// third party apis
type apis struct {
}

func New(conf *config.Config, l logger.Logger) Provider {

	bodb, err := database.Connect(conf.DB)
	if err != nil {
		l.Fatal(err)
	}
	l.Info("connected to db")

	// создаем клиента redis

	cacheClient := redis.NewClient(&redis.Options{Addr: conf.RedisAddr, Password: "", DB: 0})
	// создаем http клиента
	c := &http.Client{
		Timeout: 20 * time.Second,
	}

	return &provider{
		&resources{
			bodb,
			cacheClient,
		},
		&apis{},
		c,
		l,
	}
}

func (p *provider) GetDBConn() *sqlx.DB {
	return p.resources.db
}

func (p *provider) GetCacheClient() *redis.Client {
	return p.resources.cacheClient
}

func (p *provider) GetHTTPClient() *http.Client {
	return p.c
}

func (p *provider) Close() {
	if err := p.resources.db.Close(); err != nil {
		p.l.Error("error while closing db:", err)
	} else {
		p.l.Info("db was closed")
	}
	if err := p.resources.cacheClient.Close(); err != nil {
		p.l.Error("error while closing redis client conn:", err)
	} else {
		p.l.Info("redis client was closed")
	}

}
