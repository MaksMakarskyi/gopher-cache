package datatypes

import "fmt"

type GopherHashmap struct {
	Data map[string]string
}

func NewGopherMap() *GopherHashmap {
	return &GopherHashmap{
		make(map[string]string),
	}
}

func (gh *GopherHashmap) Hset(args []string) error {
	if len(args)%2 != 0 {
		return fmt.Errorf("ERR Odd number of arguments")
	}

	p := 0
	for p < len(args) {
		gh.Data[args[p]] = args[p+1]
		p += 2
	}

	return nil
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
		values[i] = gh.Data[key]
	}

	return values
}
