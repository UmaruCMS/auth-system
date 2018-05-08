package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-redis/redis"
	// MySQL database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type config struct {
	MySQLAddr     string `json:"mysql_addr"`
	RedisAddr     string `json:"redis_addr"`
	RedisPassword string `json:"redis_password"`
	RedisDBNumber int    `json:"redis_db_num"`
}

// Database instance
var Database *gorm.DB

// RedisClient instance
var RedisClient *redis.Client

func initDatabase(cfg *config) {
	var err error
	Database, err = gorm.Open("mysql", cfg.MySQLAddr)
	if err != nil {
		panic(err.Error())
	}
}

func initRedisClient(cfg *config) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDBNumber,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}

func getConfig() *config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err.Error())
	}
	cfg := &config{}
	json.Unmarshal(raw, cfg)
	return cfg
}

func init() {
	cfg := getConfig()
	initDatabase(cfg)
	initRedisClient(cfg)
}

// Release all resources
func Release() {
	if Database != nil {
		Database.Close()
	}
	if RedisClient != nil {
		RedisClient.Close()
	}
}
