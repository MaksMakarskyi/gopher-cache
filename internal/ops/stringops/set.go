package stringops

import (
	"fmt"

	"github.com/MaksMakarskyi/gopher-cache/internal/db"
	gobj "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/opserrors"
)

func Set(s *db.GopherDB, key string, value any) error {
	strValue, ok := value.(string)
	if !ok {
		return &opserrors.InvalidInputError{
			Operation: "SET",
			InputType: fmt.Sprintf("%T", value),
		}
	}

	obj, ok := s.Get(key)

	if !ok {
		s.Set(key, &gobj.GopherObject{
			Type: gobj.GopherString,
			Ptr:  strValue,
		})
		return nil
	}

	if obj.Type != gobj.GopherString {
		return &opserrors.WrongTypeOperationError{
			Operation: "SET",
			Type:      gobj.TypeStringMap[obj.Type],
		}
	}

	obj.Ptr = strValue
	return nil
}
