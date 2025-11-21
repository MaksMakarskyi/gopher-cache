package commandexecutor

import (
	"fmt"

	command "github.com/MaksMakarskyi/gopher-cache/internal/commands"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/queue"
)

type GopherCommandExecutor struct {
	Queue      *queue.GopherQueue
	Storage    *db.GopherDB
	CommandMap map[string]command.CommandHandler
}

func NewGopherCommandExecutor(q *queue.GopherQueue, d *db.GopherDB, cmds map[string]command.CommandHandler) *GopherCommandExecutor {
	return &GopherCommandExecutor{
		Queue:      q,
		Storage:    d,
		CommandMap: cmds,
	}
}

func (gce *GopherCommandExecutor) Start() {
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

func (gce *GopherCommandExecutor) Execute(cmd *command.GopherCommand) (string, error) {
	cmdHandler, exist := gce.CommandMap[cmd.Name]
	if !exist {
		return "", fmt.Errorf("ERR command does not exist: %s", cmd.Name)
	}

	return cmdHandler(gce.Storage, cmd.Args)
}
