package main

import "fmt"

func main()  {
	sayHelloTo("Syahrul", ".")
}

func sayHelloTo(firstName string, lasName string){
	fmt.Println("Hello", firstName, lasName)
}


// saat buat funvtion, kadang kita butuh data dari lua/ disebut parameter
// kita bisa menambahkan parameter function bisa lebih dari satu
// parameter tidakalah wajib
// jika menambahkan parameter di function, maka ketika memanggil function tersebut kita wajib memasukkan data ke parameternya



