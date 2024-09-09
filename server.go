package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"netcat/global"
)

var Connections global.Conns

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Invalid arguments\nUsage:go run . <PORT>")
		return
	}
	port := fmt.Sprint(":", args[0])
	MessageChan := make(chan global.Message)
	done := make(chan struct{})
	defer close(done)

	global.InitLogo()
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	// logs.MakeFile()
	go func(messageChan chan global.Message, done chan struct{}) {
		for {
			select {
			case msg := <-messageChan:
				Connections.SendMessage(msg, "message")
			case <-done:
				return
			}
		}
	}(MessageChan, done)
	fmt.Printf("Listening on the port %s\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(conn, MessageChan)
	}
}

func handleConnection(conn net.Conn, messageChan chan global.Message) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String())
	conn.Write(global.Logo)
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	valid, name := global.NameProcessor(&Connections, string(buffer[:len]))
	if !valid {
		_, err := fmt.Fprintf(conn, "%senter a valid unique name to connect, try something else :)%s\n", global.Colors.Red, global.Colors.Reset)
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		err := Connections.Add(global.Connection{Name: name, Conn: conn})
		if err != nil {
			fmt.Fprint(conn, err.Error())
			return
		}
		Connections.SendMessage(global.Message{Name: name, Msg: []byte(fmt.Sprintf("\n%s%s has joined our chat...%s", global.Colors.LightGray, name, global.Colors.Reset))}, "status")
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(conn, "%s[%s][%s]:%s", global.Colors.Green, now, name, global.Colors.Yellow)
	}
	for {
		len, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				Connections.Delete(name)
				Connections.SendMessage(global.Message{Name: name, Msg: []byte(fmt.Sprintf("\n%s%s has left our chat...%s", global.Colors.LightGray, name, global.Colors.Reset))}, "status")
				return
			}
			log.Fatalln(err)
		}
		if len > 0 {
			if global.IsValidMessage(string(buffer[:len])) {
				now := time.Now().Format("2006-01-02 15:04:05")
				message := global.Message{Name: name, Msg: buffer[:len], Time: now}
				messageChan <- message
				now = time.Now().Format("2006-01-02 15:04:05")
				fmt.Fprintf(conn, "%s[%s][%s]:%s", global.Colors.Green, now, name, global.Colors.Yellow)
			} else {
				now := time.Now().Format("2006-01-02 15:04:05")
				fmt.Fprintf(conn, "%s%s%s[%s][%s]:%s", global.Colors.Red, "Invalid message\n", global.Colors.Green, now, name, global.Colors.Yellow)
			}
		}
	}
}
