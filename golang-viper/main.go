package main

import (
	// "bytes"
	// "io"
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	defaults = map[string] interface{} {
		"username": "admin",
		"password": "password",
		"host": 3306,
		"database": "test",
	}

	configName = "config"
	configPaths = []string {
		".",
	}
)

type Config struct {
	Username string
	Password string
	Host string
	Port int
	Database string
}

func main() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("could not read config file: %v", err)
	}

	fmt.Println("Username from viper: ", viper.GetString("username"))
	fmt.Println("name from viper: ", viper.GetString("name"))


	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Could not decode config into struct: %v", err)
	}


	fmt.Printf("Username from struct: %s\n", config.Username)
	fmt.Printf("Config struct: %v\n", config)

	// -------------------


	changed := false 
	viper.WatchConfig()
	viper.OnConfigChange(func (e fsnotify.Event) {
		err = viper.Unmarshal(&config)
		if err != nil {
			log.Printf("Could not decode config after changed: %v\n", err)
		}
		changed = true
	})

	for !changed {
		time.Sleep(time.Second)
		fmt.Printf("\nConfig struct: %v", config)
	}
}