package main

import (
	"fmt"
	"github.com/nickdu2009/gotest/lib"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	lib.CheckErr(err)

	for _, v := range interfaces {
		fmt.Println(v)
	}
}
