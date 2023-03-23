package main

import (
	test "github.com/felix021/kitex-test/kitex_gen/kitex/felix021/test/testservice"
	"log"
)

func main() {
	svr := test.NewServer(new(TestServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
