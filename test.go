package main

import (
	"fmt"
	"time"
)

type y struct {
	T time.Time
}

func main() {
	x := &y{time.Now()}
	fmt.Println(x.T)
	time.Sleep(2 * time.Second)
	fmt.Println(x.T)

}
