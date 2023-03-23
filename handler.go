package main

import (
	"context"
	pbapi "github.com/felix021/kitex-test/kitex_gen/pbapi"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	resp = &pbapi.Response{
		Message: req.Message,
	}
	return
}
