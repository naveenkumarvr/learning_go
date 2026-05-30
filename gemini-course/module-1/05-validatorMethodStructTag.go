package main

import "fmt"

type ServerConfig struct {
	Host  string `json:"host_name"`
	port  int    `json:"port_num"`
	IsSSL bool   `json:"ssl_enabled"`
}

func (cfg *ServerConfig) checkSecurity() error {
	if !cfg.IsSSL {
		return fmt.Errorf("\nWarning: Insecure Connection")
	} else {
		return nil
	}
}

func (cfg *ServerConfig) checkPort() error {
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

	for i := range configs {
		// if err := checkPort(config.port, config.Host); err != nil {

		if err := configs[i].checkPort(); err != nil {
			fmt.Printf("\nDeployment Halted %v", err)
		} else {
			fmt.Printf("\nSUCCESS: Standard port %s, %d", configs[i].Host, configs[i].port)
		}

		if err := configs[i].checkSecurity(); err != nil {
			fmt.Printf("\nWARN: %v", err)
		} else {
			fmt.Printf("Secure Connection as SSL is set to %v", configs[i].IsSSL)
		}
	}
}
