package main

import (
	"fmt"
	"log"
	"net"
	"netcat/global"
)

func main() {
	global.InitLogo()
	ln, err := net.Listen("tcp", "localhost:2525")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Listening on the port :2525")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	// conn.Write(global.Logo)
	conn.Read(buffer)
	fmt.Println(string(buffer))
	for {
		conn.Read(buffer)
		fmt.Println(buffer)
		conn.Write(buffer)
		buffer = buffer[:0]
	}

}
