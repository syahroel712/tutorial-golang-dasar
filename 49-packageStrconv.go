package main

import (
	"fmt"
	"strconv"
)

func main(){ 
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	// strconv.Parseint("string", base(desimal = 10, binary = 2 , oktal = 8 , hexa = 16), bitSize(32, 64))
	number, err := strconv.ParseInt("100000", 2, 64)

	if err == nil {
		fmt.Println(number)
	} else{
		fmt.Println("Error", err.Error())
	}


	// strconv.FormatInt(nilai, base(desimal = 10, binary = 2 , oktal = 8 ,  hexa= 16))
	value := strconv.FormatInt(100000, 16)
	fmt.Println(value)


}


// catatan
// strcov = string conversion
// digunakan agar bisa mengkonversi data dari string ke int, boolean atau sebaliknya
// contoh
	// strconv.ParseBool(string) => string ke bool
	// strconv.ParseFloat(string) => string ke Float
	// strconv.ParseInt(string) => string ke Int
	// strconv.FormatBool(string) => Bool ke string
	// strconv.FormatFloat(string) => Float ke string
	// strconv.FormatInt(string) => Int ke string
// contoh
	// func main(){
	// 	boolean, err := strconv.ParseBool("true")

	// 	if err == nil {
	// 		fmt.Println(boolean)
	// 	} else {
	// 		fmt.Println("Error", err.Error())
	// 	}
	// }

// hampir semua strcov bisa mengembalikan err
// cara paling simple utk konversi string ke int adalah menggunakan atoi / menggunakan itoa konversi int ke string