package cmdexecutor

import (
	"fmt"

	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

type CommandDoesNotExistError struct {
	Command string
}

func (e *CommandDoesNotExistError) Error() string {
	errorMsg := fmt.Sprintf("ERR command does not exist: %s", e.Command)
	return encodingutils.FormatSimpleError(errorMsg)
}
