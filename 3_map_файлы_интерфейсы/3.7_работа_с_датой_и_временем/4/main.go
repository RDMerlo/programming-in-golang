package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	rd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strDate := strings.Trim(rd, "\n")
	t, _ := time.Parse("2006-01-02 15:04:05", strDate)

	if t.Hour() < 13 {
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	} else {
		futureDate := t.Add(time.Hour * 24)
		fmt.Println(futureDate.Format("2006-01-02 15:04:05"))
	}
}
