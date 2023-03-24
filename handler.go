package main

import (
	"context"
	echo "github.com/felix021/kitex-test/kitex_gen/echo"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *echo.EchoRequest) (resp *echo.EchoResponse, err error) {
	resp = &echo.EchoResponse{
		Message: req.Message,
	}
	return
}
