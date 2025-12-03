package setcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherset"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// member exists in set → :1\r\n
// not exists → :0\r\n
// key missing → :0\r\n
// wrong type → WRONGTYPE
// wrong args count → error

func Sismember(d *db.GopherDB, key string, member string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return encodingutils.FormatInteger(0), nil
	}

	value, ok := obj.Pointer.(*gopherset.GopherSet)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	exist := value.Sismember(member)
	return encodingutils.FormatInteger(exist), nil
}

func SismemberHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 2 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "SISMEMBER"}
	}

	return Sismember(d, args[0], args[1])
}
