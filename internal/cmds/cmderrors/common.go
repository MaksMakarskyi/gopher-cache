package cmderrors

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
	return fmt.Sprintf("wrong operation for type %s: %s", e.Type, e.Operation)
}

type InvalidInputError struct {
	Operation string
	InputType string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input of type %s for operation %s", e.InputType, e.Operation)
}
