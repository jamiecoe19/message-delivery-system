package main

import (
	"log"
	"mds/cmd/mds/server"
)

func main() {

	conn, err := server.CreateRabbitMQConnection()
	if err != nil {
		log.Fatal("rabbit connection error")
	}
	defer conn.Close()

	db, err := server.CreateSqlConnection()
	if err != nil {
		log.Fatal("db connection error")
	}
	defer db.Close()

	server.
		New().
		Listen()
}
