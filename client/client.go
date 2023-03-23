package main

import (
	"context"
	"github.com/cloudwego/kitex/transport"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/felix021/kitex-test/kitex_gen/pbapi"
	"github.com/felix021/kitex-test/kitex_gen/pbapi/echo"
)

func main() {
	cli, err := echo.NewClient(
		"hello",
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &pbapi.Request{Message: "hello world"}
		resp, err := cli.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
