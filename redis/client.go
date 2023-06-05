package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	connectionCheckInterval = time.Minute * 10
	reconnectInterval       = time.Second * 2
)

type redisClient struct {
	config           Config
	serviceName      string
	client           *redis.Client
	connectionStatus ConnectionStatus
}

func NewClient(config Config, serviceName string) *redis.Client {

	client := redisClient{
		config:      config,
		serviceName: serviceName,
	}
	ok := client.setConnection()
	if !ok {
		fmt.Println("error whole connect to redis")
	}
	go client.supportConnection()
	return client.client
}

func (r *redisClient) setConnection() (ok bool) {
	r.client = redis.NewClient(&redis.Options{
		Addr:                  fmt.Sprintf("%v:%v", r.config.Host, r.config.Port),
		Username:              r.config.User,
		Password:              r.config.Password,
		ClientName:            r.serviceName,
		ReadTimeout:           time.Duration(r.config.ReadTimeout) * time.Second,
		WriteTimeout:          time.Duration(r.config.WriteTimeout) * time.Second,
		PoolFIFO:              true,
		PoolSize:              r.config.MaxPoolSize,
		ContextTimeoutEnabled: true,
	})
	if !r.checkConnection() {
		return false
	}
	return true
}

func (r *redisClient) checkConnection() (ok bool) {
	if r.client == nil {
		return false
	}
	err := r.client.Ping(context.Background()).Err()
	if err != nil {
		r.connectionStatus.SetClosedStatus()
		return false
	}
	r.connectionStatus.SetReadyStatus()
	return true
}

func (r *redisClient) supportConnection() {
	for {
		time.Sleep(connectionCheckInterval)
		if !r.checkConnection() {
			r.reconnect()
		}
	}
}

func (r *redisClient) reconnect() {

	if r.client != nil {
		err := r.client.Close()
		if err != nil {
			os.Exit(1)
		}
	}
	for {
		ok := r.setConnection()
		if !ok {
			time.Sleep(reconnectInterval)
			continue
		}
		return
	}
}
