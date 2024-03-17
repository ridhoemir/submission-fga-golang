package main

import "submission-3/services"

func main() {
	port := ":8080"
	services.StartServer(port)
}
