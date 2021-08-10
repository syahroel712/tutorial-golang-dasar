package main

import "fmt"

func main(){
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var name = "Syahrul"
	var e = name[0]
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}



// catatan 

// di go kadang kita butuh konvesri tipe data ke tipe data lain
	// var nilai32 int32 = 32
	// var nilai64 int64 = int64(nilai32)
	// var nilai16 int16 = int16(nilai32)

// di go waktu ammil karakter,maka akan menghasilkan ke binary, maka untuk konversi ke string kembali kita bisa gunakan konversi
	// var name = "Syahrul"
	// var e = name[0]
	// var eString = string(e)
	// fmt.Println(name)
	// fmt.Println(eString)

