package opserrors

import "fmt"

type NotExistError struct {
	Key string
}

func (e *NotExistError) Error() string {
	return fmt.Sprintf("value does not exist for key: %s", e.Key)
}

type WrongTypeOperationError struct {
	Operation string
	Type      string
}

func (e *WrongTypeOperationError) Error() string {
	return fmt.Sprintf("invalid operation for type %s: %s", e.Type, e.Operation)
}

type TypeValueMismatchError struct {
	Expected string
	Got      string
}

func (e *TypeValueMismatchError) Error() string {
	return fmt.Sprintf("internal error: expected type %s, but got %s", e.Expected, e.Got)
}
