package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func main() {

	osInfo := runtime.GOOS

	fmt.Printf("Operational system name: %s\n", osInfo)


	currentUser, err := user.Current()
	if(err != nil){
		fmt.Printf("Getting user error: %v", err)
	}
	
	fmt.Printf("User name: %v\n", currentUser.Name)
	fmt.Printf("Home directory: %v\n", currentUser.HomeDir)
	

	hostname, err := os.Hostname()
	if(err != nil){
		fmt.Printf("Getting hostname: %v\n", err)
	}

	fmt.Printf("Hostname: %v\n", hostname)

	
	
}