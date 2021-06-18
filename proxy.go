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

package main

import (
	"io"
	"net"
	"strconv"
	"time"
)

type Proxy struct {
	Service *Service
}

func NewProxy(s *Service) *Proxy {
	p := new(Proxy)
	p.Service = s
	return p
}

func (p *Proxy) Run() {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(p.Service.Config.ListenPort))
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err.Error())
		}
		go func() {
			defer conn.Close()
			p.Service.ActiveConnections++
			for p.Service.State != "READY" {
				if p.Service.State == "STOPPED" || p.Service.State == "UNKNOWN" {
					p.Service.Launch()
				}
				time.Sleep(time.Second)
			}
			dest, err := net.Dial("tcp", p.Service.Config.Destination)
			if err != nil {
				p.Service.ActiveConnections--
				return
			}
			defer dest.Close()
			errsig := make(chan bool)
			go func() {
				io.Copy(dest, conn)
				errsig <- true
			}()
			go func() {
				io.Copy(conn, dest)
				errsig <- true
			}()
			<-errsig
			p.Service.ActiveConnections--
			p.Service.LastConnectionClosed = time.Now()
		}()
	}
}
