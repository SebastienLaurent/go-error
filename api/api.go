package api

import (
	"github.com/pkg/errors"

	"github.com/SebastienLaurent/go-error/subapi"
)

var ErrApiNotWorking = errors.New("api is not working")

func Test() error {
	err := subapi.Test()
	if err != nil {
		return errors.Wrapf(ErrApiNotWorking, "can't invoke api")
	}
	
	return err
}
