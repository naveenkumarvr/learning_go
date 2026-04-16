package main

import "fmt"

func main() {
	var dbHost string = "localhost"
	var dbPort int = 67700
	var isSSL bool = false

	fmt.Printf("The Db Host %s dbPort is %d and is bool %t", dbHost, dbPort, isSSL)

	// Slice
	clusterHost := []string{"prod-web-01", "", "prod-db-01"}

	for _, server := range clusterHost {
		dbHost = server
		fmt.Printf("Checking %s", dbHost)
		if dbHost == "" {
			fmt.Printf("\nERROR: Host is empty")
		}
		if dbPort > 65535 || dbPort < 1 {
			fmt.Printf("\nERROR: Invalid port")
		}
		if !isSSL {
			fmt.Printf("\nWARNING: Connecting over an insecure connection")
		}
	}

}
