package main

import "fmt"

func main(){

	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)

}

// dengan for loop
func factorialLoop(value int) int{
	result := 1
	for i :=value; i > 0; i--{
		result *= i
	}
	return result
}

// dengan recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}



// catatan recursive function
// recursive function adalah function yg memanggil function dirinya sendiri
// kadang, kita menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
// contoh kasus yg mudah diselesaikan dengan recursive function adalah Factorial
// recursive harus memiliki kondisi berhentinya