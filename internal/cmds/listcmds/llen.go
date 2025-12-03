package listcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherlist"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// list exists → return integer length: :<len>\r\n
// key missing → return :0\r\n
// wrong type → WRONGTYPE
// wrong number of arguments → error

func Llen(d *db.GopherDB, key string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return encodingutils.FormatInteger(0), nil
	}

	glist, ok := obj.Pointer.(gopherlist.GopherList)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	length := glist.Llen()
	return encodingutils.FormatInteger(length), nil
}

func LlenHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 1 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "LLEN"}
	}

	return Llen(d, args[0])
}
