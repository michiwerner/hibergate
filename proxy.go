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
