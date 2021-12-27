package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Imgur ConfigImgur
}

type ConfigImgur struct {
	ID string
}

func ParseConfig() Config {
	var config Config

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Get config dir: " + err.Error())
	}

	_, err = toml.DecodeFile(fmt.Sprintf("%s/imgur/config.toml", configDir), &config)
	if err != nil {
		log.Fatal("Decode config file: " + err.Error())
	}

	return config
}
