package main

import "fmt"

func main(){

	name := "Budiss"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Ini Bukan Eko")
	}


	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}


// catatan

// if expression
// if adalah salah satu kunci yg digunkan utk percabangan
// percabangan artinya tidak bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
// hampir di semua bahasa pemograman mendukung if expression

// else expression
// blok if dijalankan jika true
// jika salah, maka gunkan else expression

// else if expression
// kadang dalam if, kita butuh membuat beberapa kondisi
// kasus seperti ini kita bisa menggukan else if expression

// if dengan short statement
// if mendukung short statement sebelum kondisi
// hal ini sangan cocok utk membuat statement yg sederhana sebelum melakukan pengecekan terhadap kondisi
