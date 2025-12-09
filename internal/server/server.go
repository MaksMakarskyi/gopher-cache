package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmdparser"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/queue"
)

type GopherServer struct {
	Addr          string
	Listener      net.Listener
	CommandQueue  *queue.GopherQueue
	CommnadParser *cmdparser.GopherCommandParser
}

func NewGopherServer(addr string, q *queue.GopherQueue) *GopherServer {
	return &GopherServer{
		Addr:          addr,
		CommandQueue:  q,
		CommnadParser: cmdparser.NewGopherCommandParser(),
	}
}

func (gs *GopherServer) Run() error {
	ln, err := net.Listen("tcp", gs.Addr)
	if err != nil {
		return fmt.Errorf("ERR failed to run TCP server")
	}

	gs.Listener = ln
	defer ln.Close()

	log.Println("==> Gopher Cache is ready")
	log.Printf("==> Listening on %s", gs.Addr)
	log.Println("==> Server ready and accepting RESP commands (Ctrl+C to exit)")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(fmt.Errorf("error during accepting message"))
			continue
		}

		go gs.handleConnection(conn)
	}
}

func (gs *GopherServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		bytes := make([]byte, 1024)
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}

		rawCmd := string(bytes[:n])
		cmdName, cmdArgs, err := gs.CommnadParser.Parse(rawCmd)
		if err != nil {
			conn.Write([]byte(err.Error()))
		}

		responseCh := make(chan string, 1)
		cmd := cmds.NewGopherCommand(cmdName, cmdArgs, responseCh)

		gs.CommandQueue.Add(&cmd)

		response := <-responseCh
		conn.Write([]byte(response))
	}
}
