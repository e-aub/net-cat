package global

import (
	"log"
	"os"
)

var Logo []byte

func InitLogo() {
	var err error
	Logo, err = os.ReadFile("global/logo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
