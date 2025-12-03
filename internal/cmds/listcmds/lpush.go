package listcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherlist"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key missing → create list, push values to left, return list length: :<len>\r\n
// already list → push to left, return length
// wrong type → WRONGTYPE error
// wrong number of arguments → -ERR wrong number of arguments for 'LPUSH' command\r\n

func Lpush(d *db.GopherDB, key string, members []string) (string, error) {
	obj, ok := d.KVStore[key]

	var newLen int

	if !ok {
		newList := gopherlist.NewGopherList()
		newLen = newList.Lpush(members)
		d.KVStore[key] = &datatypes.GopherObject{Pointer: newList}

		return encodingutils.FormatInteger(newLen), nil
	}

	glist, ok := obj.Pointer.(gopherlist.GopherList)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	newLen = glist.Lpush(members)
	return encodingutils.FormatInteger(newLen), nil
}

func LpushHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) < 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "LPUSH"}
	}

	return Lpush(d, args[0], args[1:])
}
