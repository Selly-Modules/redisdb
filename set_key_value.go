package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// SetKeyValue ...
func SetKeyValue(ctx context.Context, key string, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		fmt.Println("redisdb - SetKeyValue error", err.Error())
		return
	}
	rdb.Set(ctx, key, b, 0)
}

// SetTTL ...
func SetTTL(ctx context.Context, key string, value interface{}, d time.Duration) {
	b, err := json.Marshal(value)
	if err != nil {
		fmt.Println("redisdb - SetTTL error", err.Error())
		return
	}
	rdb.Set(ctx, key, b, d)
}
