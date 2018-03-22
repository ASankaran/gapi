package cache

import (
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})

func WriteString(key string, value string) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		client.Del(key)
	}
}

func GetString(key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		return ""
	}
	return value
}

func DeleteString(key string) {
	client.Del(key)
}

func KeyExists(key string) bool {
	_, err := client.Get(key).Result()
	return err == nil
}
