package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	IdleSecs     int    `required:"true"`
	StopCmd      string `required:"true"`
	ReadinessCmd string `required:"true"`
	LaunchCmd    string `required:"true"`
	ListenPort   int    `required:"true"`
	Destination  string `required:"true"`
}

func NewConfig() *Config {
	c := new(Config)
	err := envconfig.Process("HIBERGATE", c)
	if err != nil {
		panic("Error processing env variables: " + err.Error())
	}
	return c
}
