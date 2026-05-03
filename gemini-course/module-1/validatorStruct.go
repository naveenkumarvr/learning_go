package main

import "fmt"


type ServerConfig struct {
	Host string
	port int
	IsSSL bool
}

func checkPort(port int, dbHost string) error {
	switch {
	case port >= 65535 || port < 1:
		return fmt.Errorf("\nERROR: Invalid port %d for host %s", port, dbHost)
	case port < 1024:
		return fmt.Errorf("\nWARN: Privileged Port %d for host %s", port, dbHost)
	default:
		return nil
	}
}
(
func main()