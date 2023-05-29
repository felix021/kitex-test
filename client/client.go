package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/rpctimeout"

	"github.com/cloudwego/kitex/client"
	"github.com/felix021/kitex-test/kitex_gen/echo"
	"github.com/felix021/kitex-test/kitex_gen/echo/echoservice"
)

func main() {
	rpctimeout.SetGlobalNeedFineGrainedErrCode()
	cli, err := echoservice.NewClient(
		"serverName",
		client.WithHostPorts("127.0.0.1:8888"),
		client.WithRPCTimeout(time.Millisecond*300),
	)
	if err != nil {
		log.Fatal(err)
	}

	timeoutByBusiness(cli)
	cancelByBusiness(cli)
	timeoutByKitex(cli)
}

func timeoutByBusiness(cli echoservice.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	req := &echo.EchoRequest{Message: "hello world"}
	resp, err := cli.Echo(ctx, req)
	processRespAndErr(resp, err)
}

func cancelByBusiness(cli echoservice.Client) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * 20)
		cancel()
		log.Println("cancel called")
	}()

	req := &echo.EchoRequest{Message: "hello world"}
	resp, err := cli.Echo(ctx, req)
	processRespAndErr(resp, err)
}

func processRespAndErr(resp *echo.EchoResponse, err error) {
	if err == nil {
		log.Println(resp)
		return
	}

	log.Print("err:", err)
	if err == kerrors.ErrCanceledByBusiness {
		log.Println("is canceled by business")
	} else if err == kerrors.ErrTimeoutByBusiness {
		log.Println("is timeout by business")
	} else if kerrors.IsTimeoutError(err) {
		log.Println("is timeout by kitex")
	} else {
		log.Fatal(err)
	}
}

func timeoutByKitex(cli echoservice.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*299)
	defer cancel()
	go func() {
		time.Sleep(time.Millisecond * 1000)
		cancel()
		log.Println("cancel called")
	}()

	req := &echo.EchoRequest{Message: "hello world"}
	resp, err := cli.Echo(ctx, req)
	processRespAndErr(resp, err)
}
