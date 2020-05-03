package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().String())
		time.Sleep(time.Second * 1)
	}
}
