package main

import (
	"os"
	"rest-api-employee/config"
	"rest-api-employee/route"
)

func main() {
	config.InitDB()
	e := route.InitRoute()
	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(":" + getPort()))
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
