package utils

import (
	"errors"
	"fmt"
)

func NewError(s string, e error) error {
	if e != nil {
		return errors.New(s + fmt.Sprintf("-> %v", e))
	}
	return errors.New(s)

}
