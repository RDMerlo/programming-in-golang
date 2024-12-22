package main

import (
	"fmt"
	"time"
)

func main() {
	var strDate string
	fmt.Scan(&strDate)

	t, _ := time.Parse(time.RFC3339, strDate)
	fmt.Println(t.Format(time.UnixDate))
}
