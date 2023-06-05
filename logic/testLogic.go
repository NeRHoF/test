package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"hash"
	test_entities "test/testEntities"
)

type TestLogic struct {
	redisRepo test_entities.RedisRepos
}

func NewTestLogic(redisRepo test_entities.RedisRepos) test_entities.TestLogic {
	return &TestLogic{
		redisRepo: redisRepo,
	}
}
func (s *TestLogic) Test1Func(ctx context.Context, request test_entities.Test1StructRequest) (test_entities.Test1StructResponse, error) {
	currentValue, err := s.redisRepo.Get(ctx, request.Key)
	if err != nil {
		return test_entities.Test1StructResponse{}, err
	}
	newValue := currentValue + request.Val
	err = s.redisRepo.Set(ctx, request.Key, newValue)
	if err != nil {
		return test_entities.Test1StructResponse{}, err
	}
	valueAfterAction, err := s.redisRepo.Get(ctx, request.Key)
	if err != nil {
		return test_entities.Test1StructResponse{}, err
	}
	response := test_entities.Test1StructResponse{Response: valueAfterAction}
	return response, nil
}

func computeHmac512(message string, secret string) string {
	key := []byte(secret)
	var h hash.Hash
	h = hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func (s *TestLogic) Test2Func(_ context.Context, request test_entities.Test2StructRequest) (test_entities.Test2StructResponse, error) {
	criptedString := computeHmac512(request.ValueToCrypt, request.ValueToCrypt)
	response := test_entities.Test2StructResponse{ResponseHMACSHA512: criptedString}
	return response, nil
}
