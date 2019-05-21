package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DB struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	DB DB `yaml:"db"`
}

var Conf *Config

func InitConfig(path string) error {

	yamlfile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read yml file error!")
		return err
	}

	c := &Config{}

	err = yaml.Unmarshal(yamlfile, c)

	if err != nil {
		log.Println("Unmarshal yaml error")
		return err
	}

	Conf = c

	return nil
}
