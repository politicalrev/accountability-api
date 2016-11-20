package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	// Load config on package init
	viper.AutomaticEnv()

	// Configure gin environment
	switch viper.GetString("environment") {
	case "development":
		gin.SetMode(gin.DebugMode)
		break
	case "test":
		gin.SetMode(gin.TestMode)
		break
	case "production":
		gin.SetMode(gin.ReleaseMode)
		break
	}
}
