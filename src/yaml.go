package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config : Yml config data
type Config struct {
	AppSecret   string `yaml:"app_secret"`
	AccessToken string `yaml:"access_token"`
	VerifyToken string `yaml:"verify_token"`
}

// Parse content yml file and return its content
func parseContentYml() string {
	content, err := ioutil.ReadFile("content.yml")
	if err != nil {
		log.Panicf("Reading file error: %s", err)
	}

	out, err := yaml.Marshal(content)
	if err != nil {
		log.Printf("Error during marshal content: %s", err)
	}

	return string(out)
}

// Read config yml file , update Config instance with this data and return content
func (c *Config) readConfigYml() *Config {
	file, err := ioutil.ReadFile("bot.config.yml")
	if err != nil {
		log.Panicf("Error opening config file: %s", err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Printf("Error during unmarshal config file: %s\n", err)
	}

	// Debug only
	fmt.Printf("content.yml: %s", c)
	return c
}

func getToken() string {
	var c Config
	c.readConfigYml()
	content, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error during json marshalling: %s\n", err)
	}

	return string(content)
}
