package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/redis-golang/config"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Pong : ", pong)

}
