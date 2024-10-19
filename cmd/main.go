package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
	"github.com/Younes-khadraoui/Error_Sentinel/utils"
	"github.com/Younes-khadraoui/Error_Sentinel/handlers"
)

func main() {
	args := os.Args
	port := utils.GetPort(args)
	app := new(internals.WebServer)
	app.GET("/", handlers.Home)
	err := app.Start(port)
	if err != nil {
		log.Panic("Error Starting The Server")
	}
	fmt.Println("Server Running on Port", port)
}
