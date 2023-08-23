package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
	MySQL  MySQL  `yaml:"mysql"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func main() {
	config := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 反序列化到结构体
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(data))
}
