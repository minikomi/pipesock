package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now())
		time.Sleep(3 * time.Second)
	}
}
