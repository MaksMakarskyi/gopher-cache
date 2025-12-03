package setcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherset"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// key missing → return :0\r\n
// exists → return number of removed members: :<count>\r\n
// wrong type → WRONGTYPE
// wrong number of arguments → error

func Srem(d *db.GopherDB, key string, members []string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return encodingutils.FormatInteger(0), nil
	}

	value, ok := obj.Pointer.(*gopherset.GopherSet)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	count := value.Srem(members)
	return encodingutils.FormatInteger(count), nil
}

func SremHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) < 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "SREM"}
	}

	return Srem(d, args[0], args[1:])
}
