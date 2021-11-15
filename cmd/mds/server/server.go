package server

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/streadway/amqp"
)

func CreateRabbitMQConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("failed to connect to rabbit")
	}
	return conn, nil
}

func CreateSqlConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/mds")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

type Server struct {
	*echo.Echo
}

func New() Server {
	var e = echo.New()
	return Server{e}
}

func (s Server) Listen() Server {
	s.Logger.Fatal(s.Start(":8087"))
	return s
}
