package main

import (
	"os/exec"
	"sync"
	"time"
)

type Service struct {
	sync.Mutex
	Config               *Config
	State                string
	ActiveConnections    int
	LastConnectionClosed time.Time
}

func NewService(c *Config) *Service {
	s := new(Service)
	s.Config = c
	s.State = "UNKNOWN"
	return s
}

func (s *Service) UpdateIdleState() {
	s.Lock()
	defer s.Unlock()
	if s.State == "READY" && s.ActiveConnections < 1 && !s.LastConnectionClosed.IsZero() && int(time.Now().Sub(s.LastConnectionClosed).Seconds()) >= s.Config.IdleSecs {
		s.State = "IDLE"
	}
}

func (s *Service) Launch() {
	s.Lock()
	defer s.Unlock()
	s.State = "LAUNCHED"
	cmd := exec.Command("/bin/sh", "-c", s.Config.LaunchCmd)
	cmd.Run()
}

func (s *Service) Stop() {
	s.Lock()
	defer s.Unlock()
	s.State = "STOPPED"
	cmd := exec.Command("/bin/sh", "-c", s.Config.StopCmd)
	cmd.Run()

}

func (s *Service) IsIdle() bool {
	return s.State == "IDLE"
}

func (s *Service) StopIfIdle() {
	if s.IsIdle() {
		s.Stop()
	}
}

func (s *Service) UpdateReadinessState() {
	s.Lock()
	defer s.Unlock()
	if s.State == "LAUNCHED" || s.State == "NOT READY" {
		cmd := exec.Command("/bin/sh", "-c", s.Config.ReadinessCmd)
		cmd.Start()
		err := cmd.Wait()
		if err == nil {
			s.State = "READY"
		} else {
			if s.State == "READY" {
				s.State = "NOT READY"
			}
		}
	}
}

func (s *Service) UpdateState() {
	s.UpdateIdleState()
	s.UpdateReadinessState()
}
