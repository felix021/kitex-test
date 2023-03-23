package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/felix021/kitex-test/kitex_gen/kitex/felix021/test"
	"github.com/felix021/kitex-test/kitex_gen/kitex/felix021/test/testservice"
)

func main() {
	cli, err := testservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &test.EchoRequest{Message: "hello world"}
		resp, err := cli.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
