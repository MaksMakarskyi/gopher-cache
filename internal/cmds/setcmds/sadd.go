package setcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherset"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key missing → create set, add members, return :<countAdded>\r\n
// exists as set → add new members, return countAdded
// elements already present do not count
// wrong type (not set) → return: -WRONGTYPE Operation against a key holding the wrong kind of value\r\n
// wrong number of arguments → return: -ERR wrong number of arguments for 'SADD' command\r\n

func Sadd(d *db.GopherDB, key string, members []string) (string, error) {
	obj, ok := d.KVStore[key]

	var count int

	if !ok {
		newSet := gopherset.NewGopherSet()
		count = newSet.Sadd(members)
		d.KVStore[key] = &datatypes.GopherObject{
			Pointer: newSet,
		}

		return encodingutils.FormatInteger(count), nil
	}

	value, ok := obj.Pointer.(*gopherset.GopherSet)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	count = value.Sadd(members)
	return encodingutils.FormatInteger(count), nil
}

func SaddHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) < 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "SADD"}
	}

	return Sadd(d, args[0], args[1:])
}
