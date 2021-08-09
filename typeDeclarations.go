package main

import "fmt"

func main(){
	type NoKtp string

	var NoKtpEko NoKtp = "12312312312"
	fmt.Println(NoKtpEko)

}



// catatan

// type declarations ialah kemampuan membuat ulang tipe data baru dari tipe data yg sidah ada / lebih mirip alias
// biasanya digunakan utk membuat alias terhadap tipe data yg sudah ada, dengan tujuan lebih mudah
	// type NoKtp string
	// type nama_alias tipe_data


