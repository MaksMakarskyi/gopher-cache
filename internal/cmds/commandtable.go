package cmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/stringcmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
)

type CommandHandler func(db *db.GopherDB, args []string) (string, error)

var GopherCommandTable map[string]CommandHandler = map[string]CommandHandler{
	"SET": stringcmds.HandleSet,
	"GET": stringcmds.GetHandler,
}
