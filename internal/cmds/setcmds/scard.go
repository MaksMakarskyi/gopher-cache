package setcmds

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherset"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// set exists → return integer count :<size>\r\n
// key missing → :0\r\n
// wrong type → WRONGTYPE
// wrong arguments → error

func Scard(d *db.GopherDB, key string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return encodingutils.FormatInteger(0), nil
	}

	value, ok := obj.Pointer.(*gopherset.GopherSet)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	count := value.Scard()
	return encodingutils.FormatInteger(count), nil
}

func ScardHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 1 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "SCARD"}
	}

	return Scard(d, args[0])
}
