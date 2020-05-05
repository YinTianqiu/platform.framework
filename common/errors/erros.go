package errors

import "fmt"

// NewError 错误
func NewError(code int) error {
	return fmt.Errorf("%v", code)
}
