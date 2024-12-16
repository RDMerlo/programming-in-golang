package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAndSplit(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error occured reading file - %v", err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	var n int
	for {
		s, err := br.ReadString(';')
		if err != nil {
			return
		}
		n++
		if s == "0;" {
			fmt.Println(n)
			return
		}
	}
}

func main() {
	ReadAndSplit("task.data")
}
