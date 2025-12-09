package app

import (
	"fmt"
	"log"

	"github.com/MaksMakarskyi/gopher-cache/internal/cliprocessor"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmdexecutor"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/queue"
	"github.com/MaksMakarskyi/gopher-cache/internal/server"
)

type App struct {
	DB          *db.GopherDB
	Queue       *queue.GopherQueue
	CmdExecutor *cmdexecutor.GopherCommandExecutor
}

func NewApp(queueSize int) *App {
	gdb := db.NewGopherDB()
	gqueue := queue.NewGopherQueue(queueSize)
	gcmdexecutor := cmdexecutor.NewGopherCommandExecutor(gqueue, gdb)

	return &App{
		gdb,
		gqueue,
		gcmdexecutor,
	}
}

func (app *App) Run(mode string, host string, port string) {
	switch mode {
	case "server":
		addr := fmt.Sprintf("%s:%s", host, port)
		gserver := server.NewGopherServer(addr, app.Queue)

		go func() {
			if err := gserver.Run(); err != nil {
				log.Fatalf("Server error: %v", err)
			}
		}()

	case "cli":
		cliProcessor := cliprocessor.NewCLIProcessor(app.Queue)

		go func() {
			if err := cliProcessor.Run(); err != nil {
				log.Fatalf("CLI Processor error: %v", err)
			}
		}()

	default:
		fmt.Println("Mode is invalid, please select one of the following options: 'server', 'cli'")
	}

	app.CmdExecutor.Start()
}
