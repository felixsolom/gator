package main

import (
	"fmt"

	"github.com/felixsolom/gator/internal/config"
)

func main() {
	configStruct, err := config.Read(".gatorconfig.json", &config.Config{})
	if err != nil {
		fmt.Println(err)
	}
	configStruct.SetUser("felix")
	config.Write(".gatorconfig.json", configStruct)
	configStruct, err = config.Read(".gatorconfig.json", &config.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", *configStruct)
}
