package global

import (
	"log"
	"os"
	"strings"
)

var Logo []byte

func InitLogo() {
	var err error
	Logo, err = os.ReadFile("global/logo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func NameProcessor(name string) (bool, string) {
	name = strings.TrimSpace(name)
	if strings.Contains(name, "\n") || len(name) == 0 {
		return false, ""
	}
	return true, name
}
