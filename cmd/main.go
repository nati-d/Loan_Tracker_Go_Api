package main

import (
	"loan_tracker/bootstrap"
	"loan_tracker/delivery/router"
	"log"
)


func main(){
	//Initialize the environment variables
	bootstrap.InitEnv()

	//print server starting
	log.Println("Server starting...")

	uri, err := bootstrap.GetEnv("MONGO_URI")
	if err != nil {
		panic(err)
	}

	client, err := bootstrap.ConnectDatabase(uri)
	if err != nil {
		panic(err)
	}

	//Call the router setup function
	router := router.SetupRouter(client)

	router.Run(":8080")



	defer bootstrap.DisconnectDatabase(client)



}