package main

import (
	"bufio"
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
	currentLine := ""
	for {
		part1 := ""
		part2 := ""

		count, err := file.Read(data)
		if err != nil {
			break
		}

		for i := 0; i < count; i++ {
			if data[i] == '\n' {
				i += 1
				part2 += string(data[i:])
				break
			} else {
				part1 += string(data[i])
			}
		}

		currentLine += part1

		if len(part2) != 0 {
			fmt.Printf("read: %s\n", currentLine)
			currentLine = part2
		}
	}
}
