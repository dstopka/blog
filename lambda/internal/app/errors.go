package app

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type PostNotFoundError struct {
	Slug string
}

func (e *PostNotFoundError) Error() string {
	return fmt.Sprintf("post with slug '%s' not found", e.Slug)
}

func (e *PostNotFoundError) Unwrap() error {
	return ErrNotFound
}

type PropertyNotFoundError struct {
	Property string
}

func (e *PropertyNotFoundError) Error() string {
	return fmt.Sprintf("property '%s' not found", e.Property)
}