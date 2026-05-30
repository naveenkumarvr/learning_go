package main

import "fmt"

func main() {
	var dbHost string = "localhost"
	var dbPort int = 5432
	var isSSL bool = false

	fmt.Printf("The Db Host %s dbPort is %d and is bool %t", dbHost, dbPort, isSSL)

	// Slice
	clusterHost := []string{"prod-web-01", "", "prod-db-01"}

	for _, server := range clusterHost {
		dbHost = server
		fmt.Printf("Checking %s", dbHost)
		// if len(dbHost) == 0 {
		// 	fmt.Printf("\nERROR: Host is empty")
		// }

		// Improvised Condtion
		if hLen := len(dbHost); hLen == 0 {
			fmt.Printf("\n Error: Host is Empyt")
		}

		if dbPort > 65535 || dbPort < 1 {
			fmt.Printf("\nERROR: Invalid port")
		} else if dbPort < 1024 {
			fmt.Printf("\n WARNING: Privileged port")
		} else {
			fmt.Printf("\n SUCCESS: Standard Port")
		}

		if !isSSL {
			fmt.Printf("\nWARNING: Connecting over an insecure connection")
		} else {
			fmt.Printf("\n [SECURE]: Connection is encrypted")
		}
	}

}
