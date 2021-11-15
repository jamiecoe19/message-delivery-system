package main

import (
	"log"
	"mds/cmd/mds/messenger"
	"mds/cmd/mds/server"
	"mds/internal"
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

	repo := internal.NewUserRepository(db)
	rabbit := internal.NewRabbitMQ(conn)

	service := messenger.NewService(repo, rabbit)

	handler := server.NewHandler(service)

	server.
		New().
		CreateRoutes(handler).
		Listen()
}
