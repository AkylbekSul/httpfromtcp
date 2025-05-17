package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 8)

	for {
		count, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Printf("read: %s\n", data[:count])
	}
}