package hashmapcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherhashmap"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key exists → return array of bulk strings, missing fields = $-1\r\n
// Example: *3\r\n$3\r\none\r\n$-1\r\n$3\r\nthree\r\n
// key missing → return array of $-1 entries
// wrong type → WRONGTYPE error
// wrong number of arguments → -ERR wrong number of arguments for 'HMGET' command\r\n

func Hmget(d *db.GopherDB, key string, fields []string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		values := make([]string, len(fields))
		for i := range values {
			values[i] = ""
		}
		return encodingutils.FormatArray(values), nil
	}

	gmap, ok := obj.Pointer.(*gopherhashmap.GopherHashmap)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	values := encodingutils.FormatArray(gmap.Hmget(fields))
	return values, nil
}

func HandleHmget(d *db.GopherDB, args []string) (string, error) {
	if len(args) <= 1 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "HMGET"}
	}

	return Hmget(d, args[0], args[1:])
}
