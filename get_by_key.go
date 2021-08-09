package redisdb

import "context"

// GetValueByKey ...
func GetValueByKey(ctx context.Context, key string) string {
	cmd := rdb.Get(ctx, key)
	if cmd == nil {
		return ""
	}
	value, _ := cmd.Result()
	return value
}
