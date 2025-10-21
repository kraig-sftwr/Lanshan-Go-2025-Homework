package main

import "fmt"

func main() {
	var n int
	fmt.Printf("请输入一个整数:")
	fmt.Scan(&n)

	fac := Factorial(n)

	fmt.Printf("该整数的阶乘为: %d", fac)

}
