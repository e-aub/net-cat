package global

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Connection struct {
	Name string
	Conn net.Conn
}

type Color struct {
	Green     string
	LightGray string
	Red       string
	Yellow    string
	Reset     string
}

var Colors = Color{Green: "\033[1;32m", LightGray: "\033[1;37m", Red: "\033[1;31m", Yellow: "\033[1;93m", Reset: "\033[1;0m"}

type Message struct {
	Name string
	Msg  []byte
	Time string
}

type Conns struct {
	Connections []Connection
	mu          sync.RWMutex
}

func (Connections *Conns) Delete(name string) {
	Connections.mu.Lock()
	for index, conn := range Connections.Connections {
		if conn.Name == name {
			Connections.Connections = append(Connections.Connections[:index], Connections.Connections[index+1:]...)
		}
	}
	Connections.mu.Unlock()
}

func (Connections *Conns) Add(connection Connection) error {
	if len(Connections.Connections) < 10 {
		Connections.mu.Lock()
		Connections.Connections = append(Connections.Connections, connection)
		Connections.mu.Unlock()
	} else {
		return fmt.Errorf("%sserver reached maximum number of connection, come back in another time :>%s", Colors.Red, Colors.Reset)
	}
	return nil
}

func (Connections *Conns) SendMessage(message Message, typ string) error {
	if typ == "message" {
		for _, conn := range Connections.Connections {
			if conn.Name != message.Name {
				now := time.Now().Format("2006-01-02 15:04:05")
				_, err := fmt.Fprintf(conn.Conn, "\n%s[%s][%s]:%s%s%s%s%s[%s][%s]:%s", Colors.Green, message.Time, message.Name, Colors.Reset, Colors.Yellow, message.Msg, Colors.Reset, Colors.Green, now, conn.Name, Colors.Yellow)
				if err != nil {
					return err
				}
			}
		}
	} else if typ == "status" {
		for _, conn := range Connections.Connections {
			if conn.Name != message.Name {
				now := time.Now().Format("2006-01-02 15:04:05")
				_, err := fmt.Fprintf(conn.Conn, "%s%s%s\n%s[%s][%s]:%s", Colors.LightGray, string(message.Msg), Colors.Reset, Colors.Green, now, conn.Name, Colors.Yellow)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
