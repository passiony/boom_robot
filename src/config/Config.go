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

type LogConfig struct {
	Log_path    string `json:"log_path" bson:"log_path"`
	Maxlines    int32  `json:"maxlines" bson:"maxlines"`
	Maxsize     int32  `json:"maxsize" bson:"maxsize"`
	Log_level   int32  `json:"log_level" bson:"log_level"`
	Log_Console int32  `json:"log_console" bson:"log_console"`
}

type Config struct {
	Server *NetConfig `json:"server" bson:"server"`
	Log    *LogConfig `json:"log" bson:"log"`
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
