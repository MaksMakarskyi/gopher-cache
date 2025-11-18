package storage

import (
	gobj "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
)

type Storage struct {
	KV_store map[string]*gobj.GopherObject
}

func NewStorage() *Storage {
	s := Storage{map[string]*gobj.GopherObject{}}
	return &s
}

func (s *Storage) Get(key string) (*gobj.GopherObject, bool) {
	obj, ok := s.KV_store[key]
	return obj, ok
}

func (s *Storage) Set(key string, value *gobj.GopherObject) {
	s.KV_store[key] = value
}
