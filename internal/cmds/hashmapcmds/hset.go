package hashmapcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherhashmap"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key does not exist → create hash + set field, return integer count of new fields: ":1\r\n"
// key exists and field is new → return: ":1\r\n"
// key exists and field overwritten → return: ":0\r\n"
// wrong type (key not hash) → -WRONGTYPE Operation against a key holding the wrong kind of value\r\n
// wrong number of arguments → -ERR wrong number of arguments for 'HSET' command\r\n

func Hset(d *db.GopherDB, key string, args []string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		newHashmap := gopherhashmap.NewGopherMap()
		count, err := newHashmap.Hset(args)
		d.KVStore[key] = &datatypes.GopherObject{
			Pointer: newHashmap,
		}

		return encodingutils.FormatInteger(count), err
	}

	ghashmap, ok := obj.Pointer.(gopherhashmap.GopherHashmap)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	count, err := ghashmap.Hset(args)
	return encodingutils.FormatInteger(count), err
}

func HandleHset(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 2 && len(args)%2 != 0 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "HSET"}
	}

	return Hset(d, args[0], args[1:])
}
