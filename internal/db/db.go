package db

import (
	dtypes "github.com/MaksMakarskyi/gopher-cache/internal/datatypes"
)

type GopherDB struct {
	KVStore map[string]*dtypes.GopherObject
}

func NewDB() *GopherDB {
	return &GopherDB{
		make(map[string]*dtypes.GopherObject),
	}
}
