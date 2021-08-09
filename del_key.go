package redisdb

import "context"

// DelKey ...
func DelKey(ctx context.Context, key string) {
	rdb.Del(ctx, key)
}