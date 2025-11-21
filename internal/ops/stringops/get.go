package stringops

import (
	"errors"
	"fmt"

	dtypes "github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/opserrors"
)

func Get(d *db.GopherDB, key string) (string, error) {
	obj, ok := d.KVStore[key]

	if !ok {
		return "", &opserrors.NotExistError{Key: key}
	}

	if obj.Type != dtypes.StringType {
		return "", &opserrors.WrongTypeOperationError{
			Operation: "GET",
			Type:      dtypes.TypeToStringMap[obj.Type],
		}
	}

	value, ok := obj.Data.(string)
	if !ok {
		return "", &opserrors.TypeValueMismatchError{
			Expected: dtypes.TypeToStringMap[dtypes.StringType],
			Got:      fmt.Sprintf("%T", obj.Data),
		}
	}

	return value, nil
}

func GetHandler(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("ERR wrong number of arguments for 'GET' command")
	}

	key := args[0]

	value, err := Get(d, key)
	if err != nil {
		return "", err
	}

	return value, nil
}
