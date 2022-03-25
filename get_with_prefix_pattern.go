package redisdb

import (
	"context"
	"fmt"
)

// GetWithPrefixPattern ...
func GetWithPrefixPattern(pattern string, limit int64) (keys []string, values []string) {
	// Init
	keys = make([]string, 0)
	values = make([]string, 0)

	var (
		ctx    = context.Background()
		cursor uint64
		err    error
	)

	keys, cursor, err = rdb.Scan(ctx, cursor, pattern, limit).Result()
	if err != nil {
		fmt.Println("err redis scan: ", err.Error())
		return
	}

	if len(keys) == 0 {
		fmt.Println("NO keys")
		return
	}

	// Get value
	for _, key := range keys {
		value := GetValueByKey(ctx, key)
		values = append(values, value)
	}

	return
}
