package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/felix021/kitex-test/kitex_gen/echo"
	"github.com/felix021/kitex-test/kitex_gen/echo/echoservice"
)

func main() {
	cli, err := echoservice.NewClient("serverName", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &echo.EchoRequest{Message: "hello world"}
		resp, err := cli.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
