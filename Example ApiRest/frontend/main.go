package main

import (
	"os"
	"app/initapp"
)

func main() {
	a := initapp.App{}
	a.Initialize()
	PORT := getPort()
	a.Run(PORT)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8005"
	}
	return ":"+port
}
