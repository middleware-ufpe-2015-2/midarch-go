package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
	"net/http"
	"shared/parameters"
	"strconv"
	"apps/calculator/implrpc"
)

func main() {

	// create new instance of calculator
	calculator := new(implrpc.Calculator)

	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Calculator", calculator)

	// associate a http handler to server
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, e := net.Listen("tcp", ":"+strconv.Itoa(parameters.CALCULATOR_PORT))
	if e != nil {
		log.Fatal("Server not started:", e)
	}

	// wait for calls
	fmt.Println("Server is running... \n")
	http.Serve(l, nil)
}


