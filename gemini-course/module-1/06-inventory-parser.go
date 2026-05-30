package main

import (
	"encoding/json"
	"fmt"
)

type ServerConfig struct {
	Host  string `json:"host_name"`
	Port  int    `json:"port_num"`
	IsSSL bool   `json:"ssl_enabled"`
}

func main() {
	rawFleetData := `[
		{"host_name": "prod-web-01", "port_num": 80, "ssl_enabled": true},
		{"host_name": "prod-db-01", "port_num": 5432, "ssl_enabled": false}
	]`

	var fleet []*ServerConfig
	if err := json.Unmarshal([]byte(rawFleetData), &fleet); err != nil {
		fmt.Printf("\nError parsing data. %v", err)
	} else {
		for i := range fleet {
			fmt.Printf("Host: %v, Port: %v, IsSSL: %v\n", fleet[i].Host, fleet[i].Port, fleet[i].IsSSL)
		}
	}
}
