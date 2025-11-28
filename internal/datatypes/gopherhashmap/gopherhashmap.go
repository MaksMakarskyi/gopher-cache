package gopherhashmap

import (
	"github.com/MaksMakarskyi/gopher-cache/internal/cmds/cmderrors"
)

type GopherHashmap struct {
	Data map[string]string
}

func NewGopherMap() *GopherHashmap {
	return &GopherHashmap{
		make(map[string]string),
	}
}

func (gh *GopherHashmap) Hset(args []string) (int, error) {
	if len(args)%2 != 0 {
		return 0, &cmderrors.WrongNumberOfArgsError{Command: "HSET"}
	}

	p := 0
	count := 0
	for p < len(args) {
		if _, ok := gh.Data[args[p]]; !ok {
			count += 1
		}
		gh.Data[args[p]] = args[p+1]
		p += 2
	}

	return count, nil
}

func (gh *GopherHashmap) Hget(key string) string {
	if value, ok := gh.Data[key]; !ok {
		return ""
	} else {
		return value
	}
}

func (gh *GopherHashmap) Hmget(keys []string) []string {
	values := make([]string, len(keys))
	for i, key := range keys {
		values[i] = gh.Hget(key)
	}

	return values
}
