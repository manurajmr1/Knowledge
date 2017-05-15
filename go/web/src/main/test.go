package main

import "time"
import "fmt"

func main() {
	// message := make(chan string, 5) // no buffer
	count := 100000000
	current := 0
	start := time.Now()

	for i := 1; i <= count; i++ {
		current += (i * 2) / (10 * (22 / 12 * (44 / 22)))
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed, current)

}
