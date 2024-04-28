package in_memory

import (
	"context"
	"errors"
	"sync"
	"time"
)

const TtlCache = 1 * time.Minute
const TickerTime = 2 * time.Second

type InMemoryCache[T any] struct {
	items map[int64]CacheValue[T]
	mxPVZ sync.RWMutex
}

type CacheValue[T any] struct {
	Item T
	TTL  time.Time
}

func NewInMemoryCache[T any]() *InMemoryCache[T] {
	cache := &InMemoryCache[T]{
		items: make(map[int64]CacheValue[T], 50),
		mxPVZ: sync.RWMutex{},
	}

	go func() {
		t := time.NewTicker(TickerTime)
		for {
			select {
			case <-t.C:
				now := time.Now()
				for key, v := range cache.items {
					if v.TTL.Before(now) {
						cache.mxPVZ.Lock()
						delete(cache.items, key)
						cache.mxPVZ.Unlock()
					}
				}
			}
		}
	}()

	return cache
}

func (c *InMemoryCache[T]) Set(_ context.Context, id int64, item T) error {
	c.mxPVZ.Lock()
	defer c.mxPVZ.Unlock()
	c.items[id] = CacheValue[T]{
		Item: item,
		TTL:  time.Now().Add(TtlCache),
	}
	return nil
}

func (c *InMemoryCache[T]) Get(_ context.Context, id int64) (T, error) {
	c.mxPVZ.RLock()
	defer c.mxPVZ.RUnlock()
	info, ok := c.items[id]
	if !ok {
		var noop T
		return noop, errors.New("cant find type by id")
	}

	return info.Item, nil
}

func (c *InMemoryCache[T]) Delete(_ context.Context, id int64) {
	c.mxPVZ.Lock()
	defer c.mxPVZ.Unlock()
	delete(c.items, id)
}
