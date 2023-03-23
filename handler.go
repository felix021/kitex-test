package main

import (
	"context"
	test "github.com/felix021/kitex-test/kitex_gen/kitex/felix021/test"
)

// TestServiceImpl implements the last service interface defined in the IDL.
type TestServiceImpl struct{}

// Echo implements the TestServiceImpl interface.
func (s *TestServiceImpl) Echo(ctx context.Context, req *test.EchoRequest) (resp *test.EchoResponse, err error) {
	resp = &test.EchoResponse{
		Message: req.Message,
	}
	return
}
