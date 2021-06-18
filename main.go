package main

import (
	"time"
)

func main() {
	config := NewConfig()
	service := NewService(config)
	proxy := NewProxy(service)
	go func() {
		for {
			service.UpdateState()
			service.StopIfIdle()
			time.Sleep(time.Second)
		}
	}()
	proxy.Run()
}
