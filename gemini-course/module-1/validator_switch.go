package main

import "fmt"

func main() {

	hostPort := map[string]int{
		"prod-web-01": 80,
		"prod-db-01":  5432,
		"bad-server":  99999,
	}

	for dbHost, port := range hostPort {
		switch {
		case port > 65534 || port < 1:
			fmt.Printf("\nERROR: Invalid port %s, %d", dbHost, port)
		case port < 1024:
			fmt.Printf("\n WARN: Privileged Port %s, %d", dbHost, port)
		default:
			fmt.Printf("\nSUCCESS: Standard port %s, %d", dbHost, port)
		}
	}
}
