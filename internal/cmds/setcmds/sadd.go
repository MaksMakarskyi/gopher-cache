package setcmds

// import (
// 	"errors"
// 	"fmt"

// 	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
// 	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
// 	"github.com/MaksMakarskyi/gopher-cache/internal/datatypes/gopherset"
// 	"github.com/MaksMakarskyi/gopher-cache/internal/db"
// )

// key missing → create set, add members, return :<countAdded>\r\n
// exists as set → add new members, return countAdded
// elements already present do not count
// wrong type (not set) → return: -WRONGTYPE Operation against a key holding the wrong kind of value\r\n
// wrong number of arguments → return: -ERR wrong number of arguments for 'SADD' command\r\n

// func Sadd(d *db.GopherDB, key string, members []string) error {
// 	obj, ok := d.KVStore[key]

// 	if !ok {
// 		newSet := gopherset.NewGopherSet()
// 		newSet.Sadd(members)
// 		d.KVStore[key] = &datatypes.GopherObject{
// 			Pointer: newSet,
// 		}

// 		return nil
// 	}

// 	if obj.Type != datatypes.SetType {
// 		return &cmderrors.WrongTypeOperationError{
// 			Operation: "GET",
// 			Type:      datatypes.TypeToStringMap[obj.Type],
// 		}
// 	}

// 	value, ok := obj.Data.(*gopherset.GopherSet)
// 	if !ok {
// 		return &cmderrors.TypeValueMismatchError{
// 			Expected: datatypes.TypeToStringMap[datatypes.StringType],
// 			Got:      fmt.Sprintf("%T", obj.Data),
// 		}
// 	}

// 	value.Sadd(members)
// 	return nil
// }

// func SaddHandler(d *db.GopherDB, args []string) (string, error) {
// 	if len(args) < 2 {
// 		return "", errors.New("ERR too few arguments for 'SADD' command")
// 	}

// 	err := Sadd(d, args[0], args[1:])
// 	if err != nil {
// 		return "", err
// 	}

// 	return "OK", nil
// }
