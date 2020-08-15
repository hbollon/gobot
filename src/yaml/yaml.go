package yaml

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

type ResponsesPool struct {
	Templates []struct {
		Messages []string `yaml:"messages"`
		Response string   `yaml:"response"`
	} `yaml:"templates"`
	DefaultResponse string `yaml:"default_response"`
}

var responsesPool ResponsesPool = parseContentYml()

// Parse content yml file and return its content
func parseContentYml() ResponsesPool {
	var pool ResponsesPool
	content, err := ioutil.ReadFile("./config/content.yml")
	if err != nil {
		log.Panicf("Reading file error: %s", err)
	}

	err = yaml.Unmarshal(content, &pool)
	if err != nil {
		log.Printf("Error during unmarshal content: %s", err)
	}

	// Debug only
	fmt.Printf("content.yml: %s\n\n", pool)
	return pool
}

// Read config yml file , update Config instance with this data and return content
func (c *Config) readConfigYml() *Config {
	file, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Panicf("Error opening config file: %s", err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Printf("Error during unmarshal config file: %s\n", err)
	}

	// Debug only
	fmt.Printf("config.yml: %s\n\n", c)
	return c
}

func GetToken() string {
	var c Config
	c.readConfigYml()
	content, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error during json marshalling: %s\n", err)
	}

	return string(content)
}
