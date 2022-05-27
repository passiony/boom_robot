package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type NetConfig struct {
	Ip   string `json:"ip" bson:"ip"`
	Port int    `json:"port" bson:"port"`
}

type Config struct {
	Server *NetConfig `json:"server" bson:"server"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func LoadConfig(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("load config with error", err)
		return
	}
	config = &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Println("json unmarshal with error ", err)
		return
	}
}
