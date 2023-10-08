package main

import (
	"encoding/json"
	"flag"
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

var (
	config Config
	path   string
)

// init 从命令行读取配置文件路径
// go run main.go -c /app/test/config.yml
func init() {
	flag.StringVar(&path, "c", "/app/config.yml", "config file path")
	// 必须设置，否则无法读取命令行参数
	flag.Parse()
	initViper()
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".") // 可以指定配置文件的绝对路径
	if path != "" {
		viper.AddConfigPath(path)
	}
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
