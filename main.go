package main

import "mygram/app/handler"

func main() {
	var PORT = ":8080"

	handler.StartServer().Run(PORT)
}
