package cache

import (
	"context"
	"encoding/json"
	"errors"
	"test/custom_errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type redisClient struct {
	redisPassword string
	redisDatabase int
	redisAddress  string
}

func (client *redisClient) GetString(key string) (string, error) {
	rdb := client.createRedisClient()
	result := rdb.Get(ctx, key)
	err := result.Err()
	if err != nil {
		if err == redis.Nil {
			return "", custom_errors.NewKeyNotFoundError(key)
		} else {
			return "", err
		}
	}

	return result.Val(), err
}

func (client *redisClient) Get(key string, object interface{}) error {
	if object == nil {
		return errors.New("object argument must not be nil")
	}

	value, err := client.GetString(key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(value), object)
	return err
}

func (client *redisClient) SetString(key, value string, duration time.Duration) error {
	rdb := client.createRedisClient()
	defer rdb.Close()
	setResult := rdb.Set(ctx, key, value, 0)
	err := setResult.Err()
	if err != nil {
		return err
	}

	// вызываем отдельно expire, потому что метод Set
	// возвращает ошибку ERR wrong number of arguments for 'set' command golang,
	// если duration подставить как последний аргумент метода Set, возможно это проблема установленного редиса
	setExpireResult := rdb.Expire(ctx, key, duration)
	return setExpireResult.Err()
}

func (client *redisClient) Set(key string, object interface{}, duration time.Duration) error {
	if object == nil {
		return errors.New("object argument must not be nil")
	}

	json, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = client.SetString(key, string(json), duration)
	return err
}

func (client *redisClient) createRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     client.redisAddress,
		Password: client.redisPassword,
		DB:       client.redisDatabase,
	})

	return rdb
}

func (client *redisClient) Delete(key string) error {
	rdb := client.createRedisClient()
	defer rdb.Close()
	result := rdb.Del(ctx, key)
	return result.Err()
}

func NewRedisClient(redisPassword string, redisDatabase int, redisAddress string) *redisClient {
	if redisAddress == "" {
		panic("redisAddress must not be empty")
	}

	return &redisClient{
		redisAddress:  redisAddress,
		redisPassword: redisPassword,
		redisDatabase: redisDatabase,
	}
}
