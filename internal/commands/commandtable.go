package command

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/stringops"
)

type CommandHandler func(db *db.GopherDB, args []string) (string, error)

var GopherCommandTable map[string]CommandHandler = map[string]CommandHandler{
	"SET": stringops.HandleSet,
	"GET": stringops.GetHandler,
}
