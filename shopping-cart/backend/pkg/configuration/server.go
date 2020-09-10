package configuration

import "fmt"

// Server specifics TCP address for the server to listen on, in the form "host:port"
type Server struct {
	// Host is the host that app run on
	Host string

	// Port is the port that app run on
	Port int
}

// GetAddress method returns TCP address
func (s Server) GetAddress() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}
