package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// вам это понадобится
const now = 1589570165
const UnixDate = "Mon Jan _2 15:04:05 MST 2006"

func main() {
	rd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strInput := strings.Trim(rd, "\n")
	arr := strings.Split(strInput, " ")
	// преобразует строку в Duration с использованием аннотаций:
	// "ns" - наносекунды,
	// "us" - микросекунды,
	// "ms" - миллисекунды,
	// "s" - секунды,
	// "m" - минуты,
	// "h" - часы.
	strDuration := arr[0] + "m" + arr[2] + "s"
	dur, err := time.ParseDuration(strDuration)
	if err != nil {
		panic(err)
	}

	curTime := time.Unix(now, 0).UTC()
	futureTime := curTime.Add(dur)

	fmt.Println(futureTime.Format(UnixDate))

}
