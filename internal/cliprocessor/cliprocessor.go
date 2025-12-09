package cliprocessor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmdparser"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/queue"
)

type CLIProcessor struct {
	CommandQueue    *queue.GopherQueue
	OutputFormatter *OutputFormatter
}

func NewCLIProcessor(q *queue.GopherQueue) *CLIProcessor {
	of := NewOutputFormatter(
		cmdparser.NewGopherCommandParser(),
	)

	return &CLIProcessor{CommandQueue: q, OutputFormatter: of}
}

func (cp *CLIProcessor) Run() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		userInput, err := reader.ReadBytes(byte('\n'))

		if err != nil {
			log.Fatal("Failed to read user input")
			return err
		}

		responseCh := make(chan string, 1)
		cmdName, cmdArgs := cp.ProcessUserInput(string(userInput))
		userCommand := cmds.NewGopherCommand(cmdName, cmdArgs, responseCh)

		cp.CommandQueue.Add(&userCommand)

		formattedOutput, err := cp.OutputFormatter.Format(<-responseCh)

		if err != nil {
			log.Fatal("Failed to format output")
			return err
		}

		fmt.Println(formattedOutput)
	}
}

func (cp *CLIProcessor) ProcessUserInput(ui string) (string, []string) {
	ui = strings.TrimSpace(ui)
	args := strings.Split(ui, " ")

	return args[0], args[1:]
}
