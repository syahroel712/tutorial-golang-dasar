package main

import "fmt"

func main(){
	total := sumAll(10, 29,91)
	fmt.Println(total)


	// slice parameter	
	slice := []int{10,10,10}
	total =  sumAll(slice...)
	fmt.Println(total)

}

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}



// catatan

// parameter yg berada di posisi terakhir dari function, memiliki kemampuan dijadikan sebuah varargs
// Varargs(variabel agrgumens) artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam array
// apa bedanya dengan parameter biasa dengan tipe data array
	// jika parameter tipe array, kita wajib utk membuat array terlebih dahulu sebelum mengirimkan ke function
	// jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma

// ciri2 variadic function adalah ada titik 3 di dalam parameternya
	// func sumAll(numbers ...int) int{
		// blok kode
	// }

// hanya bisa di parameter paling belakang dan hanya sekali
// jika menggunakan parameter variadic maka data tidak harus ada, tidak sama dengan parameter lainnya


// slice parameter
// kadang ada kasus dimana kita menggunakan Variadic Function, namaun memiliki variable beruoa slice
// kita bisa menjadikan slice sebagai vararg parameter
		// contoh
		// numbers := []int{10,10,10}
		// total := sumAll(numbers...)




