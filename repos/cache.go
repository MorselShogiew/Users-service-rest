package repos

import (
	"MorselShogiew/Users-service-rest/models"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	Set(key string, value *[]models.User) error
	Get(key string) (*[]models.User, error)
	GetClient() *redis.Client
}

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) Cache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *[]models.User) error {
	client := cache.GetClient()

	jsn, err := json.Marshal(value)
	if err != nil {
		return err
	}
	client.Set(key, jsn, cache.expires*time.Second)
	return nil
}

func (cache *redisCache) Get(key string) (*[]models.User, error) {
	client := cache.GetClient()

	value, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	var user []models.User

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
