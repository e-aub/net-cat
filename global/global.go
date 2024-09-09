package global

import (
	"log"
	"os"
	"strings"
	"sync"
)

var Logo []byte

func InitLogo() {
	var err error
	Logo, err = os.ReadFile("global/logo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func NameProcessor(connections *Conns, name string) (bool, string) {
	var mu sync.RWMutex
	name = strings.TrimSpace(name)
	if strings.Contains(name, "\n") || len(name) == 0 {
		return false, ""
	}
	mu.Lock()
	for _, conn := range connections.Connections {
		if conn.Name == name {
			return false, ""
		}
	}
	mu.Unlock()
	return true, name
}

func IsValidMessage(message string) bool {
	trimmedMessage := strings.TrimSpace(message)
	return len(trimmedMessage) > 0 && len(trimmedMessage) <= 1024
}
