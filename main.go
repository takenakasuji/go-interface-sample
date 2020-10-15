package main

import (
	"fmt"
	"github.com/takenakasuji/go-interface-sample/service"
)

func main() {
	s := service.NewConcatService("aaa", "bbb")
	r := s.DoSomething()
	fmt.Println(r)
}
