package redisdb

import (
	"context"
)

// GetWithPrefixPattern ...
func GetWithPrefixPattern(pattern string) (keys []string, values []string) {
	// Init
	keys = make([]string, 0)
	values = make([]string, 0)

	var (
		ctx    = context.Background()
		cursor uint64
		err    error
	)

	keys, cursor, err = rdb.Scan(ctx, cursor, pattern, 1000000).Result()
	if err != nil {
		return
	}

	if len(keys) == 0 {
		return
	}

	// Get value
	for _, key := range keys {
		value := GetValueByKey(ctx, key)
		values = append(values, value)
	}

	return
}
