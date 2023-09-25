package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	now = time.Now()
	var year int = now.Year()
	month := now.Month()
	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
