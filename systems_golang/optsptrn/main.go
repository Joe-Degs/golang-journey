package main

import (
    "fmt"
    "log"
)

func main() {
    server, err := NewServer(
	   Proto("tcp"),
	   HostPort("[::]:34423"),
    )
    if err != nil {
	   log.Fatal(err)
    }
    fmt.Println(server.proto, server.address())
    if server.Run(); err != nil {
	   log.Fatal(err)
    }
}
