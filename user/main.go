package main

import (
	"log"

	"github.com/ssibrahimbas/go-uuid/config"
	"github.com/ssibrahimbas/go-uuid/database"
	userModule "github.com/ssibrahimbas/go-uuid/internal"
	"github.com/ssibrahimbas/go-uuid/server"
)

func main() {
	c := config.New()
	databases := database.New(c)
	databases.ConnectAll()
	srv := server.New()
	um := userModule.New(databases.Postgres.Db, srv.App)
	um.Init()
	log.Fatal(srv.Listen(c.Server.Host + ":" + c.Server.Port))
}
