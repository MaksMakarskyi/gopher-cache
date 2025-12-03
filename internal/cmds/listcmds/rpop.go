package listcmds

import (
	"strconv"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherlist"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

// list exists and non-empty → pop from right
// key missing or list empty → return: "$-1\r\n"
// wrong type → WRONGTYPE error
// wrong number of arguments → -ERR wrong number of arguments for 'LPOP' command\r\n

func Rpop(d *db.GopherDB, key string, count int) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return encodingutils.GetNullBulkString(), nil
	}

	glist, ok := obj.Pointer.(*gopherlist.GopherList)
	if !ok {
		return "", &cmderrors.WrongTypeOperationError{}
	}

	if glist.Llen() == 0 {
		return encodingutils.GetNullBulkString(), nil
	}

	deletedItems := glist.Rpop(count)
	return encodingutils.FormatArray(deletedItems), nil
}

func RpopHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 2 || len(args) != 1 {
		return "", &cmderrors.WrongNumberOfArgsError{Command: "RPOP"}
	}

	if len(args) == 1 {
		return Rpop(d, args[0], 1)
	}

	count, err := strconv.Atoi(args[1])
	if err != nil {
		return "", &cmderrors.ValueNotIntegerError{}
	}

	return Rpop(d, args[0], count)
}
