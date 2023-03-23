package main

import (
	pbapi "github.com/felix021/kitex-test/kitex_gen/pbapi/echo"
	"log"
)

func main() {
	svr := pbapi.NewServer(
		new(EchoImpl),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
