// Code generated by Kitex v0.4.4. DO NOT EDIT.

package testservice

import (
	server "github.com/cloudwego/kitex/server"
	test "github.com/felix021/kitex-test/kitex_gen/kitex/felix021/test"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler test.TestService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}