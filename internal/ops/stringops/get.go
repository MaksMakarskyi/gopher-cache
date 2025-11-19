package stringops

import (
	"fmt"

	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	gobj "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/opserrors"
)

func Get(s *db.GopherDB, key string) (string, error) {
	obj, ok := s.Get(key)

	if !ok {
		return "", &opserrors.NotExistError{Key: key}
	}

	if obj.Type != gobj.GopherString {
		return "", &opserrors.WrongTypeOperationError{
			Operation: "GET",
			Type:      gobj.TypeStringMap[obj.Type],
		}
	}

	value, ok := obj.Ptr.(string)
	if !ok {
		return "", &opserrors.TypeValueMismatchError{
			Expected: gobj.TypeStringMap[gobj.GopherString],
			Got:      fmt.Sprintf("%T", obj.Ptr),
		}
	}

	return value, nil
}
