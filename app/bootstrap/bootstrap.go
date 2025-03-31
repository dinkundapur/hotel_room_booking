package bootstrap

import (
	"context"
	"hotel-room-booking/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// All the connections can be initialized here
func Boot(ctx context.Context) bool {

	//Initialize rooms
	utils.InitializeRooms()

	if strings.ToLower(viper.GetString("APP_ENV")) == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	return true
}

// Close all conections opened in the above Boot() functions
func Sleep(ctx context.Context) bool {
	//Code to close connections if any

	return true
}
