package logs

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func MakeFile() {
// 	content, err := os.ReadFile("logs/logFileNames.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	AsciifileNames := strings.Split(string(content), "\n")
// 	aId := strings.TrimPrefix(AsciifileNames[len(AsciifileNames)-1], "conversation")
// 	id, err := atoi(aId)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "err while creating logs file")
// 		return

// 	}
// 	toCreateName := fmt.Sprintf("logs/conversation%d", id+1)
// 	os.Create(toCreateName)
// 	file, err := os.OpenFile("logs/logFileNames.txt", os.O_APPEND|os.O_WRONLY, 0o664)
// 	if err != nil{
// 		log.Fatalln(err)
// 	}
// 	file.WriteString("\nconversation")

// func atoi(str string) (int, error) {
// 	if len(str) == 0 {
// 		return 0, errors.New("empty string")
// 	}

// 	start := 0

// 	result := 0
// 	for i := start; i < len(str); i++ {
// 		char := str[i]
// 		if char < '0' || char > '9' {
// 			return 0, fmt.Errorf("invalid character: %c", char)
// 		}
// 		result = result*10 + int(char-'0')
// 	}

// 	return result, nil
// }
