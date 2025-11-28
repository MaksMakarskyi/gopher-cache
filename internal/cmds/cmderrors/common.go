package cmderrors

import (
	"fmt"

	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

type WrongNumberOfArgsError struct {
	Command string
}

func (e *WrongNumberOfArgsError) Error() string {
	errorMsg := fmt.Sprintf("ERR wrong number of arguments for '%s' command", e.Command)
	return encodingutils.FormatSimpleErrors(errorMsg)
}

type WrongTypeOperationError struct{}

func (e *WrongTypeOperationError) Error() string {
	errorMsg := "WRONGTYPE Operation against a key holding the wrong kind of value"
	return encodingutils.FormatSimpleErrors(errorMsg)
}
