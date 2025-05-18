package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("messages.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			fmt.Print("read: ", line)
		}
		if err != nil {
			break
		}
	}
}
