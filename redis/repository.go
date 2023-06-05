package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	test_entities "test/testEntities"
)

type testRedisRepo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) test_entities.RedisRepos {
	return testRedisRepo{client: client}
}

func (p testRedisRepo) Get(ctx context.Context, key string) (int64, error) {
	value, err := p.client.Get(
		ctx,
		key,
	).Int64()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (p testRedisRepo) Set(ctx context.Context, key string, value int64) error {
	resp := p.client.Set(
		ctx,
		key,
		value,
		0,
	)
	if resp.Err() != nil {
		return resp.Err()
	}
	return nil
}
