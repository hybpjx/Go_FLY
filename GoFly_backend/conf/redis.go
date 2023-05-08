package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var redisClient *redis.Client
var DEFAULT_DURATION = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       0,
	})
	// ping下看是否有错误
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

// Set 设置值
func (r *RedisClient) Set(key string, value any, rest ...any) error {

	redisDuration := DEFAULT_DURATION
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			redisDuration = v
		}
	}

	return redisClient.Set(context.Background(), key, value, redisDuration).Err()
}

// Get 取值
func (r *RedisClient) Get(key string) (any, error) {
	return redisClient.Get(context.Background(), key).Result()
}

// Delete 删除值
func (r *RedisClient) Delete(key ...string) error {
	return redisClient.Del(context.Background(), key...).Err()
}

// GetExpireDuration 获取redis 过期时长
func (r *RedisClient) GetExpireDuration(key string) (time.Duration, error) {
	return redisClient.TTL(context.Background(), key).Result()
}
