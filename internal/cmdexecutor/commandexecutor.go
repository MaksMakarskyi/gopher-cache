package cmdexecutor

import (
	"log"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/queue"
)

type GopherCommandExecutor struct {
	Queue      *queue.GopherQueue
	Storage    *db.GopherDB
	CommandMap map[string]cmds.CommandHandler
}

func NewGopherCommandExecutor(q *queue.GopherQueue, d *db.GopherDB) *GopherCommandExecutor {
	return &GopherCommandExecutor{
		Queue:      q,
		Storage:    d,
		CommandMap: cmds.GopherCommandTable,
	}
}

func (gce *GopherCommandExecutor) Start() {
	log.Print("started command executor")

	for cmd := range gce.Queue.WaitForCommands() {
		result, err := gce.Execute(cmd)
		if err != nil {
			cmd.ResponseCh <- err.Error()
		} else {
			cmd.ResponseCh <- result
		}

		close(cmd.ResponseCh)
	}
}

func (gce *GopherCommandExecutor) Execute(cmd *cmds.GopherCommand) (string, error) {
	cmdHandler, exist := gce.CommandMap[cmd.Name]
	if !exist {
		return "", &CommandDoesNotExistError{Command: cmd.Name}
	}

	return cmdHandler(gce.Storage, cmd.Args)
}
