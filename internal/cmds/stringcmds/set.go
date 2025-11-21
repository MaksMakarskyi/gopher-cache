package stringcmds

import (
	"errors"
	"fmt"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
	dtypes "github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
	"github.com/MaksMakarskyi/gopher-cache/internal/db"
)

func Set(d *db.GopherDB, key string, value any) error {
	strValue, ok := value.(string)
	if !ok {
		return &cmderrors.InvalidInputError{
			Operation: "SET",
			InputType: fmt.Sprintf("%T", value),
		}
	}

	obj, ok := d.KVStore[key]

	if !ok {
		d.KVStore[key] = &dtypes.GopherObject{
			Type: dtypes.StringType,
			Data: dtypes.NewGopherString(strValue),
		}
		return nil
	}

	obj.Data = dtypes.NewGopherString(strValue)
	return nil
}

func HandleSet(d *db.GopherDB, args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("ERR wrong number of arguments for 'SET' command")
	}

	key := args[0]
	value := args[1]

	err := Set(d, key, value)
	if err != nil {
		return "", err
	}

	return "OK", nil
}
