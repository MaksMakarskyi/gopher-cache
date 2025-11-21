package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/MaksMakarskyi/gopher-cache/internal/db"
)

type GopherServer struct {
	Addr         string
	Listener     net.Listener
	CommandQueue chan<- string
}

func NewGopherServer(db *db.GopherDB, addr string, queue chan<- string) *GopherServer {
	return &GopherServer{
		Addr:         addr,
		CommandQueue: queue,
	}
}

func (gs *GopherServer) Run() error {
	ln, err := net.Listen("tcp", gs.Addr)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to run TCP server"))
		return fmt.Errorf("failed to run TCP server")
	}

	log.Printf("started listening %s", gs.Addr)

	gs.Listener = ln
	defer ln.Close()

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
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}

		fmt.Printf("request: %s", bytes)
		line := fmt.Sprintf("Echo: %s", bytes)
		fmt.Printf("response: %s", line)

		gs.CommandQueue <- line
	}
}
