package redisdb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/logrusorgru/aurora"
)

var (
	rdb *redis.Client
)

// Connect ...
func Connect(uri, password string) error {
	ctx := context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
		DB:       0, // use default DB
	})

	// Ping
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error when connect to Redis", uri, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + uri))

	return nil
}






