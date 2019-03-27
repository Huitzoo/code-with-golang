package main

import (
	"os"
	"app/initapp"
)

func main() {
	//instance a struct type App, this struct is in app/initapp
	a := initapp.App{}
	// THIS PART IS FOR CONNET TO MY LOCAL DATABASE
	a.Initialize("vive","vive","cyza")
	PORT := getPort()
	a.Run(PORT)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
		//log.Fatal("$PORT must be set")
	}
	return ":"+port
}
