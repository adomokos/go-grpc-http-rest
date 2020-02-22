package main

import (
	"fmt"
	"os"

	"github.com/adomokos/go-grpc-http-rest-microservice-tutorial/pkg/cmd"
)

func main() {
	fmt.Print("hello")
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
