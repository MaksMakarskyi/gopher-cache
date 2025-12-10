package cmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/hashmapcmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/listcmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/pingcmd"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/setcmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/stringcmds"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
)

type CommandHandler func(db *db.GopherDB, args []string) (string, error)

var GopherCommandTable map[string]CommandHandler = map[string]CommandHandler{
	"PING":      pingcmd.HandlePing,
	"SET":       stringcmds.HandleSet,
	"GET":       stringcmds.GetHandler,
	"HSET":      hashmapcmds.HandleHset,
	"HGET":      hashmapcmds.HandleHget,
	"HMGET":     hashmapcmds.HandleHmget,
	"SADD":      setcmds.SaddHandler,
	"SREM":      setcmds.SremHandler,
	"SISMEMBER": setcmds.SismemberHandler,
	"SCARD":     setcmds.ScardHandler,
	"LPOP":      listcmds.LpopHandler,
	"RPOP":      listcmds.RpopHandler,
	"LPUSH":     listcmds.LpushHandler,
	"RPUSH":     listcmds.RpushHandler,
	"LLEN":      listcmds.LlenHandler,
}
