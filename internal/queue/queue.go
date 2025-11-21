package queue

import command "github.com/MaksMakarskyi/gopher-cache/internal/commands"

type GopherQueue struct {
	CommandQueueCh chan *command.GopherCommand
}

func NewGopherQueue(bufferSize int) *GopherQueue {
	return &GopherQueue{
		CommandQueueCh: make(chan *command.GopherCommand, bufferSize),
	}
}

func (gq *GopherQueue) Add(c *command.GopherCommand) {
	gq.CommandQueueCh <- c
}

func (gq *GopherQueue) WaitForCommands() <-chan *command.GopherCommand {
	return gq.CommandQueueCh
}
