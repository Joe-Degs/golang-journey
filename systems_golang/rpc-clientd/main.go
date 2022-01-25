package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	buf := bufio.NewReader(os.Stdin)
	for {
		cmd, err := buf.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf(cmd)
	}
}
