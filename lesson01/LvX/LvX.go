package main

import "fmt"

func main() {
	var n int
	sum := 0
	count := 0
	for i := 0; ; i++ {
		fmt.Printf("Please input an integer:")
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		sum = sum + n
		count++
	}

	aver := Average(sum, count)

	if aver >= 60 {
		fmt.Printf("Average: %f, pass", aver)
	} else {
		fmt.Printf("Average: %f, failed", aver)
	}
}
