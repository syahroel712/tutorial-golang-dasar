package main

import (
	"belajar-golang-dasar/packageimport"
	"fmt"
)

func main() {
	// bisa
	packageimport.SayHello("Syahrul")
	fmt.Println(packageimport.Application)
	
	// tidak bisa
	// packageimport.sayGoodBye("Syahrul")
	// fmt.Println(packageimport.version)
}



// catatan

// access modifier
// Di bahasa pemograman lain,ada kata kunci yg bisa digunakan utk menentukan access modifier(hak akses terhadap suatu function, varaible atau semacamnya) terhadap suatu function atau variable
// di go, utk menentukan access modfier, cukup dengan nama function/variablenya
// jika Diawali Huruf Besar maka bisa diakses dari package lain
	// contoh
	// var Application = "Golang"
	// func SayGoodBye(name string) string {
	// 	return name
	// }

// jika Diawali Huruf Kecil maka tidak bisa diakses dari package lain
	// contoh
	// var version = "1.0.1"
	// func sayGoodBye(name string) string {
	// 	return name
	// } 


// contoh pada file 43-package&import.go dan packageimport/helper.go
	// jika dijalankan akan berhasil
	// jika kita ganti SayHello yg di helper.go memnjadi sayHello maka tidak akan berhasil karena akses modifier