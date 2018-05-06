package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/redis-golang/config"
)

type User struct {
	UserID    int    `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserType  string `json:"user_type"`
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	r := mux.NewRouter()

	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Test data"))
	})

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		// if error stop process
		if err != nil {
			fmt.Printf("Error during parse User model : %v", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Here can be your logic inserting data to DB or etc.

		// Logic check if user olready exists in redis DB
		userKey := fmt.Sprintf("userid_%d", user.UserID)

		_, err = client.Get(userKey).Result()
		if err == redis.Nil {
			fmt.Println(userKey + " key does not exist")

			// userKey not exists int this case will do insert data to Redis
			err = client.Set(userKey, user, time.Minute).Err()
			if err != nil {
				fmt.Println("Error during insert data to redis : " + err.Error())
			}
		}

		// In case if user was inserted in DB.
		w.WriteHeader(http.StatusCreated)

	}).Methods("POST")

	sr := &http.Server{
		Handler: r,
		Addr:    "localhost:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(sr.ListenAndServe())
}
