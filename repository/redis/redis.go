package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	_interface "github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository/interface"
)

const (
	errRedisNil = "redis: nil"
)

type redisClient struct {
	RedisClient *redis.Client
}

// NewMicroservices will create an object that represent services related to this project
func NewRedis(redisHost string, redisport string, redispassword string) (_interface.RedisCache, error) {
	//client := redis.NewFailoverClient(&redis.FailoverOptions{MasterName: redisMaster, SentinelAddrs: chHost})
	return &redisClient{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisHost, redisport),
			Password: redispassword,
		}),
	}, nil
}

// Close is used for closing the redis connection
func (r *redisClient) Close() error {
	if r.RedisClient != nil {
		if err := r.RedisClient.Close(); err != nil {
			return err
		}
		r.RedisClient = nil
	}
	return nil
}
