package redisdb

import (
	"context"
	"encoding/json"
)

// GetValueByKey ...
func GetValueByKey(ctx context.Context, key string) string {
	cmd := rdb.Get(ctx, key)
	if cmd == nil {
		return ""
	}
	value, _ := cmd.Result()
	return value
}

// GetJSON ...
func GetJSON(ctx context.Context, key string, result interface{}) (ok bool) {
	v := GetValueByKey(ctx, key)
	if v == "" {
		return false
	}
	if err := json.Unmarshal([]byte(v), result); err != nil {
		return false
	}
	return true
}
