package stringcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherstring"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key does not exist → store string, return: +OK\r\n
// key exists → overwrite value, return: +OK\r\n
// wrong number of arguments → return: -ERR wrong number of arguments for 'SET' command\r\n

func Set(d *db.GopherDB, key string, value string) string {
	obj, ok := d.KVStore[key]

	if !ok {
		d.KVStore[key] = &datatypes.GopherObject{
			Pointer: gopherstring.NewGopherString(value),
		}
	}

	obj.Pointer = gopherstring.NewGopherString(value)
	return encodingutils.FormatSimpleString("OK")
}

func HandleSet(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "SET"}
	}

	return Set(d, args[0], args[1]), nil
}
