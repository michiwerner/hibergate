/*
Copyright 2020 Michael Werner

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hibergate

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
