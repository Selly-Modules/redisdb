package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
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
