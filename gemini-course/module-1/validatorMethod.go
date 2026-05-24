package main

import "fmt"

type ServerConfig struct {
	Host  string
	port  int
	IsSSL bool
}

func (cfg ServerConfig) checkSecurity() error {
	if !cfg.IsSSL {
		return fmt.Errorf("\nWarning: Insecure Connection")
	} else {
		return nil
	}
}

func (cfg ServerConfig) checkPort() error {
	switch {
	case cfg.port > 65535 || cfg.port < 1:
		return fmt.Errorf("\nERROR: Invalid port %d for host %s", cfg.port, cfg.Host)
	case cfg.port < 1024:
		return fmt.Errorf("\nWARN: Privileged Port %d for host %s", cfg.port, cfg.Host)
	default:
		return nil
	}
}

func main() {
	configs := []ServerConfig{
		{Host: "prod-web-01", port: 80, IsSSL: true},
		{Host: "prod-db-01", port: 5432, IsSSL: false},
		{Host: "bad-server", port: 99999, IsSSL: true},
	}

	for _, config := range configs {
		// if err := checkPort(config.port, config.Host); err != nil {

		if err := config.checkPort(); err != nil {
			fmt.Printf("\nDeployment Halted %v", err)
		} else {
			fmt.Printf("\nSUCCESS: Standard port %s, %d", config.Host, config.port)
		}

		if err := config.checkSecurity(); err != nil {
			fmt.Printf("\nWARN: %v", err)
		} else {
			fmt.Printf("Secure Connection as SSL is set to %v", config.IsSSL)
		}
	}
}
