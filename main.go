package main

import (
	"rest-api-employee/config"
	"rest-api-employee/route"
)

func main() {
	config.InitDB()
	e := route.InitRoute()
	e.Logger.Fatal(e.Start(":8080"))
}
