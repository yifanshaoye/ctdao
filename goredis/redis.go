package goredis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// dsn: user:password@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func GetRedisInstance(host, port, password string) (*redis.Client, error) {
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "6379"
	}
	addr := host + ":" + port
	client := redis.NewClient(&redis.Options{
		Addr:               addr,
		Password:           password,
		DB: 0,
		PoolSize: 500,
		IdleTimeout: time.Second,
		IdleCheckFrequency: 10 * time.Second,
		MinIdleConns: 3,
		MaxRetries: 2,
		DialTimeout: 2 * time.Second,
	})

	for i := 0; i < 3; i++ {
		if pong, err := client.Ping().Result(); err != nil {
			fmt.Printf("connect redis %s, error: %+v\n", addr + " " + password, err)
			continue
		} else {
			fmt.Println("connect redis ok: ", pong)
			return client, nil
		}
	}

	return client, errors.New("redis connect fails")
}