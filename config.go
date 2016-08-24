package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Source *XMPPClient `yaml:source`
	Destination *XMPPClient `yaml:destination`
}

type XMPPClient struct {
	JID string `yaml:jid`
	Password string `yaml:password`
	Endpoint string `yaml:endpoint`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}