package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)

	i := 0
	for {
		<-ticker.C
		fmt.Println("i=", i)
		i++
		if i == 5 {
			ticker.Stop()
			break
		}
	}
}
