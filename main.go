package main

import (
	"context"
	"fmt"
	route_v1 "hotel-room-booking/app/routes/v1"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Creating Background context to receive interruption
	// Create a context that will be cancelled when an interrupt signal is received
	ctx, cancel := context.WithCancel(context.Background())

	// Loading environment variables
	settingEnvironment()

	app := gin.New()
	api := app.Group("/api")
	route_v1.Route(&gin.Context{}, api)

	// Run on PORT 8089
	server := &http.Server{
		Addr:    ":" + viper.GetString("HTTP_PORT"),
		Handler: app,
	}

	// Listen for interrupt signals and cancel the context when one is received
	go terminate(ctx, cancel, server)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server error: %v\n", err)
		cancel()
	}
	fmt.Println("Server Running", server.Addr)
}

// Loading Environment Variables
func settingEnvironment() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func terminate(ctx context.Context, cancel context.CancelFunc, server *http.Server) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	log.Print("Received interrupt signal, shutting down hotel-room-booking-service...")

	//clearing connection and closing server
	closeServer(ctx, server)
}

func closeServer(ctx context.Context, server *http.Server) {
	// Server will be closing here.
	server.Shutdown(ctx)
}
