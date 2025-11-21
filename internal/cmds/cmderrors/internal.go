package cmderrors

import "fmt"

type TypeValueMismatchError struct {
	Expected string
	Got      string
}

func (e *TypeValueMismatchError) Error() string {
	return fmt.Sprintf("internal error: expected type %s, but got %s", e.Expected, e.Got)
}
