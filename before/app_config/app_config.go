package app_config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type AppConfig struct {
	DB string `yaml:"db"`
}

var Cfg AppConfig

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)

	if err != nil {
		fmt.Println(err)
	}
}
