package main

import (
	"log"

	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	gopherdb := db.NewDB()
	commandqueue := make(chan string)
	app := server.NewGopherServer(gopherdb, "localhost:8080", commandqueue)
	app.Run()
}
