package cache

import "github.com/go-redis/redis"

func Connect(options []*redis.Options) map[*redis.Client]struct{} {
	clients := make(map[*redis.Client]struct{})
	for i := range options {
		clients[redis.NewClient(options[i])] = struct{}{}
	}
	return clients
}
