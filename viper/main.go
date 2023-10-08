package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
	MySQL  MySQL  `yaml:"mysql"`
	Array  Array  `yaml:"array"`
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

type Array struct {
	Names map[string][]any `yaml:"names"`
	Ages  []int            `yaml:"ages"`
}

var config Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".") // 可以指定配置文件的绝对路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 反序列化到结构体
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func main() {
	// watch 配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
		log.Println("config changed:", config)
	})

	data, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(data))
	time.Sleep(1000 * time.Second)
}
