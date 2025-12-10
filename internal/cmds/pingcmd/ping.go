package pingcmd

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

func HandlePing(d *db.GopherDB, args []string) (string, error) {
	if len(args) > 1 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "PING"}
	}

	if len(args) == 1 {
		return encodingutils.FormatSimpleString(args[0]), nil
	}

	return encodingutils.FormatSimpleString("PONG"), nil
}
