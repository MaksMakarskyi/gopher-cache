package stringops

import (
	"fmt"

	gobj "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/opserrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/storage"
)

func Get(s *storage.Storage, key string) (string, error) {
	obj, ok := s.Get(key)

	// Check if value exist
	if !ok {
		return "", &opserrors.NotExistError{Key: key}
	}

	// Check type
	if obj.Type != gobj.GopherString {
		return "", &opserrors.WrongTypeOperationError{
			Operation: "GET",
			Type:      gobj.TypeStringMap[obj.Type],
		}
	}

	// Ensure type match
	value, ok := obj.Ptr.(string)
	if !ok {
		return "", &opserrors.TypeValueMismatchError{
			Expected: gobj.TypeStringMap[gobj.GopherString],
			Got:      fmt.Sprintf("%T", obj.Ptr),
		}
	}

	return value, nil
}
