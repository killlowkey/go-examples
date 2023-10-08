package main

import "log"

type Person struct {
	name string
	age  int
}

var p *Person

func init() {
	p = &Person{name: "ray"}
}

type Config struct {
	p *Person
}

func NewConfig(p *Person) *Config {
	return &Config{
		p: p,
	}
}

// main 多个地方共用一个指针，原先位置指针修改后，其他地方也变了
func main() {
	log.Printf("%+v\n", p)
	log.Println("after changed")
	NewConfig(p).p.age = 100
	log.Printf("outside %+v\n", p)
	log.Printf("inner %+v\n", p)
}
