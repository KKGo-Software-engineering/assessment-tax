package config

import "fmt"

func (s *ServerConfig) HTTPAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
