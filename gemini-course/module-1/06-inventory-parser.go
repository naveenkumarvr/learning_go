package main

import (
	"encoding/json"
	"fmt"
)

type ServerConfig struct {
	Host  string `json:"host_name"`
	Port  int    `json:"Port_num"`
	IsSSL bool   `json:"ssl_enabled"`
}

func (cfg *ServerConfig) checkPort() error {
	switch {
	case cfg.Port > 65535 || cfg.Port < 1:
		return fmt.Errorf("\nERROR: Invalid Port %d for Host %s", cfg.Port, cfg.Host)
	case cfg.Port < 1024:
		return fmt.Errorf("\nWARN: Privileged Port %d for Host %s", cfg.Port, cfg.Host)
	default:
		return nil
	}
}

func (cfg *ServerConfig) checkSecurity() error {
	if !cfg.IsSSL {
		return fmt.Errorf("\nWarning: Insecure Connection")
	} else {
		return nil
	}
}

func main() {
	rawFleetData := `[
		{"host_name": "prod-web-01", "Port_num": 80, "ssl_enabled": true},
		{"host_name": "prod-db-01", "Port_num": 5432, "ssl_enabled": false}
	]`

	var fleet []*ServerConfig
	if err := json.Unmarshal([]byte(rawFleetData), &fleet); err != nil {
		fmt.Printf("\nError parsing data. %v", err)
	} else {
		for i := range fleet {
			fmt.Printf("\nHost: %v, Port: %v, IsSSL: %v", fleet[i].Host, fleet[i].Port, fleet[i].IsSSL)

			// CheckPort
			if err := fleet[i].checkPort(); err != nil {
				fmt.Printf("\nDeployment Halted %s", err.Error())
			} else {
				fmt.Printf("\nSuccess: Standard Port: %s, %d", fleet[i].Host, fleet[i].Port)
			}

			//SSL Check
			if err := fleet[i].checkSecurity(); err != nil {
				fmt.Printf("%v", err)
			} else {
				fmt.Printf("\nSecure Connection as SSL is set to %v", fleet[i].IsSSL)
			}
		}
	}
}
