package stringcmds

import (
	"errors"
	"fmt"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	dtypes "github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherstring"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
)

func Get(d *db.GopherDB, key string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return "", &cmderrors.NotExistError{Key: key}
	}

	if obj.Type != dtypes.StringType {
		return "", &cmderrors.WrongTypeOperationError{
			Operation: "GET",
			Type:      dtypes.TypeToStringMap[obj.Type],
		}
	}

	value, ok := obj.Data.(*gopherstring.GopherString)
	if !ok {
		return "", &cmderrors.TypeValueMismatchError{
			Expected: dtypes.TypeToStringMap[dtypes.StringType],
			Got:      fmt.Sprintf("%T", obj.Data),
		}
	}

	return value.Get(), nil
}

func GetHandler(d *db.GopherDB, args []any) (string, error) {
	strArgs, err := ExpectStrings(args)
	if err != nil {
		return "", err
	}
	if len(args) != 1 {
		return "", errors.New("ERR wrong number of arguments for 'GET' command")
	}

	value, err := Get(d, strArgs[0])
	if err != nil {
		return "", err
	}

	return value, nil
}
