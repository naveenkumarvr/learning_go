package main

import "fmt"

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

func main() {

	hostPort := map[string]int{
		"prod-web-01": 80,
		"prod-db-01":  5432,
		"bad-server":  99999,
	}

	for dbHost, port := range hostPort {
		if err := checkPort(port, dbHost); err != nil {
			fmt.Printf("\nDeployment Halted %v", err)
		} else {
			fmt.Printf("\nSUCCESS: Standard port %s, %d", dbHost, port)
		}
	}
}
