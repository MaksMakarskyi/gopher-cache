package queue

import "github.com/MaksMakarskyi/gopher-cache/internal/cmds"

type GopherQueue struct {
	CommandQueueCh chan *cmds.GopherCommand
}

func NewGopherQueue(bufferSize int) *GopherQueue {
	return &GopherQueue{
		CommandQueueCh: make(chan *cmds.GopherCommand, bufferSize),
	}
}

func (gq *GopherQueue) Add(c *cmds.GopherCommand) {
	gq.CommandQueueCh <- c
}

func (gq *GopherQueue) WaitForCommands() <-chan *cmds.GopherCommand {
	return gq.CommandQueueCh
}
