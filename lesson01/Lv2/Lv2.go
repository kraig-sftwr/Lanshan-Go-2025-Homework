package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 1001; i++ {
		sum = sum + i
	}

	fmt.Printf("1到1000的和为: %d", sum)
}
