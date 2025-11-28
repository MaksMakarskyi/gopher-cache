package stringcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherstring"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key exists and value is string → return bulk string: "$<len>\r\n<value>\r\n"
// key does not exist → return null bulk: "$-1\r\n"
// wrong type (not string) → return: -WRONGTYPE Operation against a key holding the wrong kind of value\r\n
// wrong number of arguments → return: -ERR wrong number of arguments for 'GET' command\r\n

func Get(d *db.GopherDB, key string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return "$-1\r\n", nil
	}

	value, ok := obj.Pointer.(*gopherstring.GopherString)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	str := value.Get()
	return encodingutils.FormatBulkString(str), nil
}

func GetHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 1 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "GET"}
	}

	return Get(d, args[0])
}
