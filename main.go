package main

import (
	"fmt"
	"test/configuration"
	"test/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration.LoadConfig("config.yml")
	r := gin.Default()
	fmt.Println("Running server!")
	registry.RegisterServicesAndRoutes(r)
	r.Run(fmt.Sprintf(":%v", configuration.C.ServerPort))
}
