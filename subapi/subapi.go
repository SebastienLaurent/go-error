package subapi

import "errors"

var ErrCantProcess = errors.New("can't process subapi")

func Test() error {
	return ErrCantProcess
}