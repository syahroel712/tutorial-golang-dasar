package main

import "fmt"

func main(){

	name := "Ansis"

	switch name {
		case "Ani" :
			fmt.Println("Hello Ani")
		case "Eko" :
			fmt.Println("Hello Eko")
		default :
			fmt.Println("Nama tidak ada")
	}


	// switch length := len(name); length > 5 {
	// 	case true:
	// 		fmt.Println("Nama terlalu panjang")
	// 	case false:
	// 		fmt.Println("Nama sudah benar")
	// }


	length := len(name)
	switch {
		case length > 10 :
			fmt.Println("Nama terlalu panjang")
		case length > 5 :
			fmt.Println("Nama lumayan panjang")
		default :
			fmt.Println("Nama sudah benar")
	}
}


// catatan

// selain if, utk percabangan kita juga bisa gunkan switch
// switch sanagt sederhana dibandingkan if
// biasanya switch digunakan utk melakukan pengecekan ke kondisi dalam satu variable
// sama dengan if, maka switch juga punya short statement sebelum varible yg akan di cek kondisinya

// switch tanpa kondisi
// kondisi di switch tidak wajib
// jika tidak menggunakan kondisi di switch, kita bisa menambahkan kondisi tersebut di setiap case