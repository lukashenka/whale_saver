package main

import (
	"./config"
	"fmt"
	"os"
)
func main()  {
	config := config.Configuration{}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	config.Load(path+"/.env.dist/config.yaml")
	fmt.Println(config.GetVolumes())
}