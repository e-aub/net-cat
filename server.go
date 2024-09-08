package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"netcat/global"
)

type Connection struct {
	Name string
	Conn net.Conn
}

type Conns struct {
	Connections []Connection
	mu          sync.Mutex
}

var Connections Conns

func (Connections *Conns) Delete(name string) {
	Connections.mu.Lock()
	for index, conn := range Connections.Connections {
		if conn.Name == name {
			Connections.Connections = append(Connections.Connections[:index], Connections.Connections[index+1:]...)
		}
	}
	Connections.mu.Unlock()
}

func (Connections *Conns) Add(Connection Connection) {
	Connections.mu.Lock()
	Connections.Connections = append(Connections.Connections, Connection)
	Connections.mu.Unlock()
}

func main() {
	global.InitLogo()
	ln, err := net.Listen("tcp", ":2525")
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
	defer conn.Close()
	conn.Write(global.Logo)
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	valid, name := global.NameProcessor(string(buffer[:len]))
	if !valid {
		_, err := fmt.Fprint(conn, "\033[31menter a valid name to connect\033[0m\n")
		if err != nil {
			log.Fatalln(err)
		}
		return
	} else {
		for _, connection := range Connections.Connections {
			_, err := connection.Conn.Write([]byte(fmt.Sprintf("%s has joined our chat...\n", name)))
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		Connections.Add(Connection{Name: name, Conn: conn})
		for {
			len, err := conn.Read(buffer)
			if err != nil {
				if err.Error() == "EOF" {
					Connections.Delete(name)
					return
				}
				log.Fatalln(err)
			}
			if len > 0 {
				fmt.Println(string(buffer[:len]))
			}

		}
	}
}
