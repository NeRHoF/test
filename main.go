package main

import (
	"fmt"
	"os"
	"test/http"
	"test/logic"
	"test/redis"

	"github.com/labstack/echo"
)

func main() {
	serviceName := "testService"
	redisConfig := redis.Config{
		Host:         "127.0.0.1",
		Port:         "6379",
		User:         "redisUser",
		Password:     "1234",
		ReadTimeout:  10,
		WriteTimeout: 10,
		MaxPoolSize:  10,
	}

	var (
		redisClient = redis.NewClient(redisConfig, serviceName)
		redisRepo   = redis.NewRedisRepository(redisClient)
		testLogic   = logic.NewTestLogic(redisRepo)
	)
	//HTTP
	e := echo.New()
	http.NewTestHandler(e.Group(""), testLogic)
	err := e.Start(":6666")
	if err != nil {
		fmt.Printf("error while starting http, err : %s", err)
		os.Exit(2)
	}

	select {}
}
