package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		li := sc.Text()
		fmt.Println(li)
	}
	defer conn.Close()
	fmt.Println("\n Hi")
}
