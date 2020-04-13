package errors

import (
	"fmt"
)

type ErrResourceNotFound struct {
	name string
}

func (e *ErrResourceNotFound) Error() string {
	return fmt.Sprintf("resource %s not found", e.name)
}

func NewErrResourceNotFound(resource string) error {
	return &ErrResourceNotFound{}
}

type ErrUnauthorize struct {
}

func (e *ErrUnauthorize) Error() string {
	return fmt.Sprintf("Unauthorize")
}

func NewErrUnauthorize() error {
	return &ErrUnauthorize{}
}
