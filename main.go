package main

import (
	echo "github.com/felix021/kitex-test/kitex_gen/echo/echoservice"
	"log"
)

func main() {
	svr := echo.NewServer(new(EchoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
