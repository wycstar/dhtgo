package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var CONFIG Config

type MongoConfig struct {
	Host       string `yaml:"host"`
	Port       int16  `yaml:"port"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

type DHTConfig struct {
	Bootstrap []string `yaml:"bootstrap"`
}

type Config struct {
	Mongo MongoConfig `yaml:"mongodb"`
	DHT   DHTConfig   `yaml:"dht"`
}

func (c *Config) Init(path string) *Config {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, c)
	if err != nil {
		panic(err)
	}
	return c
}
