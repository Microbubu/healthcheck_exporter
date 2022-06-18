package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Domain  string       `yaml:"domain"`
	LogFile string       `yaml:"logFile"`
	Tasks   []CallerTask `yaml:"tasks"`
}

type CallerTask struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Url        string `yaml:"url"`
	HttpMethod string `yaml:"httpMethod"`
	Interval   int    `yaml:"interval.seconds"`
}

func ReadYamlConfig() (*YamlConfig, error) {
	config := &YamlConfig{}
	if f, err := os.Open("config.yaml"); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(config)
	}
	return config, nil
}
