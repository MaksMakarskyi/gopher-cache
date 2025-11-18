package stringops

import (
	gobj "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/opserrors"
	"github.com/MaksMakarskyi/gopher-cache/internal/storage"
)

func Set(s *storage.Storage, key string, value string) error {
	obj, ok := s.Get(key)

	// Create new if does not exist
	if !ok {
		s.Set(key, &gobj.GopherObject{
			Type: gobj.GopherString,
			Ptr:  value,
		})
		return nil
	}

	// Check type
	if obj.Type != gobj.GopherString {
		return &opserrors.WrongTypeOperationError{
			Operation: "SET",
			Type:      gobj.TypeStringMap[obj.Type],
		}
	}

	// Set value
	obj.Ptr = value
	return nil
}
