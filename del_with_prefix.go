package redisdb

import (
	"context"
	"fmt"
)

// DelWithPrefix ...
func DelWithPrefix(ctx context.Context, prefix string) {
	iter := rdb.Scan(ctx, 0, prefix, 0).Iterator()
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			fmt.Println("Cannot delete redis key with prefix", prefix)
		}
	}
	if err := iter.Err(); err != nil {
		fmt.Println("Cannot (LOOP) delete redis key with prefix", prefix)
	}
}
