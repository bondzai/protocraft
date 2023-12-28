package main

import (
	"github.com/bondzai/protocraft/config"
	"github.com/bondzai/protocraft/database"
	"github.com/bondzai/protocraft/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
