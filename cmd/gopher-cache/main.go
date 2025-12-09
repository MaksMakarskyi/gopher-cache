package main

import (
	"flag"
	"log"

	"github.com/MaksMakarskyi/gopher-cache/internal/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mode := flag.String("mode", "server", "Run mode: 'server' listens for TCP connections, 'cli' reads commands from stdin")
	host := flag.String("host", "localhost", "Address to bind the TCP server to")
	port := flag.String("port", "6379", "TCP port the server listens on")
	queueSize := flag.Int("queueSize", 100, "Maximum number of commands buffered in the command queue")

	flag.Parse()

	app := app.NewApp(*queueSize)
	app.Run(*mode, *host, *port)
}
