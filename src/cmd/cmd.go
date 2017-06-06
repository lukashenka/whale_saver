package main

import (
	"config"
	"fmt"
	"os"
)
func main()  {
	configuration := config.Configuration{}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	configuration.Load(path+"/.env.test/config.yaml")
	fmt.Println(configuration.GetVolumes())
}