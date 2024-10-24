package main

import (
	"shortLink/database"
	"shortLink/routes"
)

func main(){
	database.StartDB()
	routes.HandleRequest()
}