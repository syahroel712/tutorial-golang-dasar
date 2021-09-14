package main

import (
	"os"
	"fmt"
)


func main(){
	// os.Args
		// args := os.Args
		// fmt.Println(args)
		// fmt.Println(args[1])
		// fmt.Println(args[2])

	// os.Hostname
		name, err := os.Hostname()

		if err == nil {
			fmt.Println("Hostname : " ,name)
		} else {
			fmt.Println("Error : ", err)
		}

}


// catatan
// Package os
// Go telah menyediakan banyak sekali package bawaan, contohny package os
// package os berisikan fungsionalitas utk mengakses fitur sistem operasi secara independen(bisa digunakan disemua sistem operasi)
// golang.org/pgk/os/

// os.Args
	// ketika perintah ini dijalankan
		// 	func main(){
			// 	args := os.Args
			// 	fmt.Println(args)
		// }
	// maka hasilnya adalah lokasi filenya seperti ini 
		// [C:\Users\SWIMOC\AppData\Local\Temp\go-build932404648\b001\exe\46-packageOs.exe]
		
	// jika perntah menjalakan nya ditambahkan argument seperti berikut
		// go run 46-packageOs.go syahrul 712 
		
	// maka hasilnya adalah lokasi filenya seperti ini 
		// [C:\Users\SWIMOC\AppData\Local\Temp\go-build932404648\b001\exe\46-packageOs.exe syahrul 712]

	// jika ingin akses argument ke2 sampai seterusnya
		// fmt.Println(args[1])
		// fmt.Println(args[2])
		// fmt.Println(args[3])

// os.Hostname
	// digunakan utk mengambil nama host dari sistem operasi yg dipakai

// Getenv 
	// undtuk ambil evironmetn di pc