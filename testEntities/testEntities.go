package test_entities

import (
	"context"
)

type Test1StructRequest struct {
	Key string `json:"key"`
	Val int64  `json:"val"`
}
type Test1StructResponse struct {
	Response int64 `json:"test"`
}

type Test2StructRequest struct {
	SecretKey    string `json:"key"`
	ValueToCrypt string `json:"s"`
}
type Test2StructResponse struct {
	ResponseHMACSHA512 string `json:"responseHMACSHA512"`
}

type TestLogic interface {
	Test1Func(ctx context.Context, request Test1StructRequest) (Test1StructResponse, error)
	Test2Func(_ context.Context, request Test2StructRequest) (Test2StructResponse, error)
}
type TcpSendRpc interface {
}
type RedisRepos interface {
	Get(ctx context.Context, key string) (int64, error)
	Set(ctx context.Context, key string, value int64) error
}
