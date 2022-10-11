package repos

import (
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/models"
	"MorselShogiew/Users-service-rest/provider"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	SetUsersList(value *[]models.User)
	GetUsersList() *[]models.User
}

type redisCache struct {
	c *redis.Client
	l logger.Logger
}

func NewRedisCache(p provider.Provider, l logger.Logger) Cache {
	return &redisCache{p.GetCacheClient(), l}
}

func (cache *redisCache) SetUsersList(value *[]models.User) {
	jsn, err := json.Marshal(value)
	if err != nil {
		return
	}
	if err := cache.c.Set("users:list", jsn, 60*time.Second).Err(); err != nil {
		cache.l.Error("error on set users list:", err)
	}
}

func (cache *redisCache) GetUsersList() *[]models.User {

	value, err := cache.c.Get("users:list").Result()
	if err != nil {
		return nil
	}

	var user []models.User

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return nil
	}

	return &user
}
