package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	host := flag.String("host", "localhost", "host for the Gopher server")
	port := flag.String("port", "6379", "port for the Gopher server") // 6379 - default Redis port

	flag.Parse()
	addr := fmt.Sprintf("%s:%s", *host, *port)

	gopherdb := db.NewDB()
	commandqueue := make(chan string)
	app := server.NewGopherServer(gopherdb, addr, commandqueue)
	app.Run()
}
