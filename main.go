package main

import (
	"github.com/gen95mis/short-url/internal/database"
	"github.com/gen95mis/short-url/internal/service"
	"github.com/gen95mis/short-url/internal/transport"
)

func main() {

	conn, err := database.NewConn()
	if err != nil {
		panic(err)
	}

	db, err := database.New(conn)
	if err != nil {
		panic(err)
	}

	service := service.New(db)

	if err := transport.Service(service); err != nil {
		panic(err)
	}
}
