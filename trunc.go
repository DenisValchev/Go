package main

import "fmt"

func Trunc() {
	var x float64
	fmt.Scan(&x)
	var y int
	y = int(x)
	fmt.Printf("%d", y)
}
