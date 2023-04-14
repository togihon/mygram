package main

import (
	"mygram/app/handler"
	"os"
)

func main() {
	//created by togihon
	var PORT = os.Getenv("PORT")

	handler.StartServer().Run(":" + PORT)
}
