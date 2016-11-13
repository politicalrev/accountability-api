package main

import (
	"log"

	"github.com/politicalrev/accountability-api/ui"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	ui.SetupRoutes(router)

	log.Printf("Running in %s mode\n", viper.GetString("environment"))
	log.Fatal(router.Run(":" + viper.GetString("port")))
}
