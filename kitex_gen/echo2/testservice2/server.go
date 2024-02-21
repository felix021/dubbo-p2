// Code generated by Kitex v0.8.0. DO NOT EDIT.
package testservice2

import (
	server "github.com/cloudwego/kitex/server"
	echo2 "github.com/felix021/dubbo-p2/kitex_gen/echo2"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler echo2.TestService2, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler echo2.TestService2, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}