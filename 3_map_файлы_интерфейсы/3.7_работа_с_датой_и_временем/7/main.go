package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	var past, future time.Time
	//13.03.2018 14:00:15,12.03.2018 14:00:15
	rd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	strAll := strings.Trim(rd, "\n")
	arr := strings.Split(strAll, ",")
	t0, _ := time.Parse("02.01.2006 15:04:05", arr[0])
	t1, _ := time.Parse("02.01.2006 15:04:05", arr[1])

	if t0.Before(t1) {
		past = t0
		future = t1
	} else {
		past = t1
		future = t0
	}

	difference := future.Sub(past)
	fmt.Println(difference)
}
