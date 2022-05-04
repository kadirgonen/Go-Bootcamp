package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

// go mod init github.com/kadirgonen/Go-Bootcamp/lesson_3
// go get go get github.com/google/uuid
func main() {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.String())
}
