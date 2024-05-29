package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nameless9000/cm2workshop/server"
)

func main() {
	var port uint16 = 8962

	if len(os.Args) > 1 {
		if new_port, err := strconv.ParseUint(os.Args[1], 10, 16); err == nil {
			port = uint16(new_port)
		} else {
			fmt.Println("Usage: cm2workshop [port: uint16]")
			os.Exit(1)
		}
	}

	server.StartServer(port)
}
