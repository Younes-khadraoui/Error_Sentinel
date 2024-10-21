package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Younes-khadraoui/Error_Sentinel/handlers"
	"github.com/Younes-khadraoui/Error_Sentinel/internals"
	"github.com/Younes-khadraoui/Error_Sentinel/middleware"
	"github.com/Younes-khadraoui/Error_Sentinel/utils"
)

func main() {
	args := os.Args
	port := utils.GetPort(args)
	app := internals.NewWebServer()
	app.GET("/", handlers.Home)
	app.GET("/retry", middleware.PreventCrash(handlers.Retry))
	app.GET("/panic", middleware.PreventCrash(handlers.Panic))
	app.GET("/error", handlers.Error)
	app.GET("/health", handlers.Health)

	fmt.Println("Server Running on Port", port)
	err := app.Start(port)
	if err != nil {
		log.Panic("Error Starting The Server")
	}
}
