package main

import (
	"github.com/reinaldodribeiro/api-movon/database"
	"github.com/reinaldodribeiro/api-movon/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}