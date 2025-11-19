package db

import (
	gobj "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
)

type GopherDB struct {
	KVStore map[string]*gobj.GopherObject
}

func NewDB() *GopherDB {
	s := GopherDB{map[string]*gobj.GopherObject{}}
	return &s
}

func (s *GopherDB) Get(key string) (*gobj.GopherObject, bool) {
	obj, ok := s.KVStore[key]
	return obj, ok
}

func (s *GopherDB) Set(key string, value *gobj.GopherObject) {
	s.KVStore[key] = value
}
