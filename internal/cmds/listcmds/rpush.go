package listcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherlist"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key missing → create list, append, return list length
// already list → append, return length
// wrong type → WRONGTYPE error
// wrong number of arguments → -ERR wrong number of arguments for 'RPUSH' command\r\n

func Rpush(d *db.GopherDB, key string, members []string) (string, error) {
	obj, ok := d.KVStore[key]

	var newLen int

	if !ok {
		newList := gopherlist.NewGopherList()
		newLen = newList.Rpush(members)
		d.KVStore[key] = &datatypes.GopherObject{Pointer: newList}

		return encodingutils.FormatInteger(newLen), nil
	}

	glist, ok := obj.Pointer.(*gopherlist.GopherList)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	newLen = glist.Rpush(members)
	return encodingutils.FormatInteger(newLen), nil
}

func RpushHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) < 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "RPUSH"}
	}

	return Rpush(d, args[0], args[1:])
}
