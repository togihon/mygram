package main

import (
	"mygram/app/handler"
	"os"
)

func main() {
	var PORT = os.Getenv("PORT")

	handler.StartServer().Run(":" + PORT)
}
