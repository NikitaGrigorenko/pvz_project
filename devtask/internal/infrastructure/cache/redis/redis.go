package redis

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis[T any] struct {
	client *redis.Client
}

const TTLRedis = time.Minute * 10

func NewRedis[T any](opt *redis.Options) *Redis[T] {
	client := redis.NewClient(opt)
	return &Redis[T]{
		client,
	}
}

func (r *Redis[T]) Set(ctx context.Context, id int64, info T) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, strconv.FormatInt(id, 10), data, TTLRedis).Err()
}

func (r *Redis[T]) Get(ctx context.Context, id int64) (T, error) {
	res := r.client.Get(ctx, strconv.FormatInt(id, 10))
	if res.Val() != "" {
		var data T
		_ = json.Unmarshal([]byte(res.Val()), &data)
		return data, nil
	}
	var noop T
	return noop, errors.New("no such id in redis")
}

func (r *Redis[T]) Delete(ctx context.Context, id int64) {
	r.client.Del(ctx, strconv.FormatInt(id, 10))
}
