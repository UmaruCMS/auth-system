package config

import (
	"encoding/json"
	"io/ioutil"
	// MySQL database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type config struct {
	MySQLAddr string `json:"mysql_addr"`
}

// Database instance
var Database *gorm.DB

func initDatabase(cfg *config) {
	var err error
	Database, err = gorm.Open("mysql", cfg.MySQLAddr)
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
}

// Release all resources
func Release() {
	if Database != nil {
		Database.Close()
	}
}
