package main

import "fmt"

func main() {
	sayHelloWithFilter("Syahrul", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

// tanpa type declarations
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

// dengan type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}


func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}



// catatan function as parameter
// function tidak hanya bisa disimpan di dalam variable sebagai value
// namun juga bisa digunakan sebagai paramater utk function lain


// function type declaration
// kdaadng jiaka fucntion terlalu panjang, agak ribet utk menuliskannya di dalam parameter
// Type Declarations juga bisa digunakan utk membuat alisa functin sehinga akan mempermudah menggunakan function parameter