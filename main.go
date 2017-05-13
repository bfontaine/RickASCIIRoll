package main

import (
	"fmt"
	"time"
)

func main() {
	for _, frame := range frames {
		fmt.Print("\033[;f")
		fmt.Println(frame)
		time.Sleep(100 * time.Millisecond)
	}
}
