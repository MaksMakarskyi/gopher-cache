package hashmapcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherhashmap"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// field exists → return bulk string: "$<len>\r\n<value>\r\n"
// field missing → return: "$-1\r\n"
// key missing → return: "$-1\r\n"
// wrong type → -WRONGTYPE Operation against a key holding the wrong kind of value\r\n
// wrong number of arguments → -ERR wrong number of arguments for 'HGET' command\r\n

func Hget(d *db.GopherDB, key string, field string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return encodingutils.GetNullBulkString(), nil
	}

	ghashmap, ok := obj.Pointer.(gopherhashmap.GopherHashmap)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	str := ghashmap.Hget(field)
	return encodingutils.FormatBulkString(str), nil
}

func HandleHget(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "HGET"}
	}

	return Hget(d, args[0], args[1])
}
