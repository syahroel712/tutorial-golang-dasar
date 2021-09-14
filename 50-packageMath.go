package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.4))
	fmt.Println(math.Round(1.5))
	
	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Ceil(1.5))

	fmt.Println(math.Max(1.4, 10))
	fmt.Println(math.Min(1, 10))
}


// catatan
// package yg berisikan constant dan fungsi matematika
// contoh
	// math.Round(float64)
	// math.Floor(float64)
	// math.Ceil(float64)
	// math.Max(float64)
	// math.Min(float64)