package main

import (
	"cproject/pkg/db"
	"fmt"
	"log"
	"net/http"
	"sync"

	"cproject/internal/config"

	"github.com/spf13/viper"
)

var once sync.Once

func loadConfig(key string) *config.Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	//value, ok := viper.Get(key).(string)
	conf := &config.Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Fatal("Invlaid type assertion")
	}

	return conf
}
func main() {
	config := loadConfig("DB_CONN_STRING")
	// TODO pass connection variableshere
	database := db.NewConnection(config.DSN)

	server := NewServer(":3000", database, config)
	err := http.ListenAndServe(":3001", server.Handler())
	if err != nil {
		fmt.Print(err)
	}
}
