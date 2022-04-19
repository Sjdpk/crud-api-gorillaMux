package main

import (
	"golang-crud/database"
	"golang-crud/routes"
)

func main() {
	database.Connect()
	routes.RoutingHandler()
}
