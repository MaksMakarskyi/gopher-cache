package cmdparser

import "fmt"

func ExpectStrings(args []any) ([]string, error) {
	out := make([]string, len(args))
	for i, arg := range args {
		s, ok := arg.(string)
		if !ok {
			return nil, fmt.Errorf("ERR argument %d must be a bulk string", i)
		}
		out[i] = s
	}
	return out, nil
}
