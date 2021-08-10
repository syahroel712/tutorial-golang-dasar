package main

import "fmt"

func main(){
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma lima = ", 3.5)
}



// catatan
// di golang manganut case sensitive

// ada 2 jenis tipe data number di go
	// 1. Integer	
			// pembagian integer yang biasa
			// * . int8 (-128  === 127) 
			// * . int16 (-32768 === 32767)
			// * . int32 (-21247483647 === 21247483647)
			// * . int64 (-9223372036854775808 === 9223372036854775807)
			// pembagian integer yang unsigned
			// * . uint8 (0 === 255)
			// * . uint16 (0 === 65528)
			// * . uint32 (0 === 4294967295)
			// * . uint64 (0 === 18446744073709551615)

	// 2. Floating Point (desimal)
			// pembaian 
			// float32 (1.18*10^38  === 3.4*10^38)
			// float64 (2.23*10^308  === 1.80*10^308)
			// complex64
			// complex128

// Alias di dalam integer
	// byte => unit8
	// rune => int32
	// int => tergantung sistem operasi (kalau 32 bit maka int32 / 64 bit maka int64)
	// uint => tergantung sistem operasi (kalau 32 bit maka uint32 / 64 bit maka uint64)