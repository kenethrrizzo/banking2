package main

import (
	"fmt"

	"github.com/kenethrrizzo/banking2/config"
)

func main() {
	// Load config
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Use config
	fmt.Println(config.DB.Host)
	fmt.Println(config.DB.Port)
	fmt.Println(config.DB.User)
	fmt.Println(config.DB.Password)
	fmt.Println(config.Server.Port)
}
