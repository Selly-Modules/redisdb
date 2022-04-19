package redisdb

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

// NewMutex ...
func NewMutex(name string) *Mutex {
	pool := goredis.NewPool(rdb)

	rs := redsync.New(pool)

	opts := []redsync.Option{
		redsync.WithRetryDelay(time.Millisecond * 100), // retry after 100 ms
		redsync.WithTries(10),                          // try 10 times
	}

	m := rs.NewMutex(name, opts...)

	return &Mutex{mu: m}
}

// Mutex ...
type Mutex struct {
	mu *redsync.Mutex
}

// Lock ...
func (m Mutex) Lock() error {
	return m.mu.Lock()
}

// Unlock ...
func (m Mutex) Unlock() (bool, error) {
	return m.mu.Unlock()
}
