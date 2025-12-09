package main

import (
	"flag"
	"log"

	"github.com/MaksMakarskyi/gopher-cache/internal/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mode := flag.String("mode", "cli", "Mode: 'cli' | 'server'. If cli is chosen, user should input commands into the terminal. If server wis chosen, user should send commands over TCP connection")
	host := flag.String("host", "localhost", "host for the Gopher server")
	port := flag.String("port", "6379", "port for the Gopher server")
	queueSize := flag.Int("queueSize", 100, "size of the command queue")

	flag.Parse()

	app := app.NewApp(*queueSize)
	app.Run(*mode, *host, *port)
}
