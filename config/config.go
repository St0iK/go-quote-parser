package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Configuration []ParserConfig `yaml:"parser_configuration"`
	QuotesFolder  string         `yaml:"QuotesFolder"`
}

type ParserConfig struct {
	Type      string `yaml:"Type"`
	Author    string `yaml:"Author"`
	QuoteText string `yaml:"QuoteText"`
	File      string `yaml:"File"`
	Category  string `yaml:"Category,omitempty"`
	Tags      string `Tags:"Tags,omitempty"`
}

func (c *Configuration) GetConf() *Configuration {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c Configuration) GetConfForFile(filename string) ParserConfig {
	log.Printf("Getting configuration for %s", filename)
	for _, v := range c.Configuration {
		if v.File == c.QuotesFolder+"/"+filename {
			return v
		}
	}

	return ParserConfig{}
}
