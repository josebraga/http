package main

import (
	"fmt"

	"github.com/josebraga/http/httptls"
)

func main() {
	s := httptls.NewTLSServer()
	c := httptls.NewTLSClient()

	fmt.Println(s)
	fmt.Println(c)
}
