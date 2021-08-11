package main

import "fmt"

func main(){

	// pass by value
		// address1 := Address{"Subang", "Jabar", "Indo"}
		// address2 := address1

		// address2.City = "Bandung" // data di address1 city subang tidak akan berubah

		// fmt.Println(address1)
		// fmt.Println(address2)

	// pointer
		// address1 := Address{"Subang", "Jabar", "Indo"}
		// address2 := &address1 // proses pointer ke address 1

		// address2.City = "Bandung" // data di address1 city subang juga berubah

		// fmt.Println(address1)
		// fmt.Println(address2)
		
	// operator *
		// address1 := Address{"Subang", "Jabar", "Indo"}
		// address2 := &address1 // proses pointer ke address 1

		// address2.City = "Bandung" // data di address1 city subang juga berubah
		
		// fmt.Println(address1)

		// // address2 = &Address{"Jakarta", "Dki Jakarta", "Indonesia"} // hal ini membuat data baru tanpa mengubah data yg lama
		// // fmt.Println(address2)
		
		// *address2 = Address{"Jakarta", "Dki Jakarta", "Indonesia"} // hal ini juga akan merubah data yang mengacu pada Address
		// fmt.Println(address2)

	// function new
		var address1 Address = Address{"subang", "jabar", "indo"}
		var address2 *Address = new(Address)

		fmt.Println(address1)
		fmt.Println(address2)
		
		address2.City = "Padang"
		fmt.Println(address2)



}


type Address struct {
	City, Province, Country string
}





// catatan


// pass by value

	// secara default di go semua varialbe itu passing by value, bukan by referance
	// artinay, jika kita mengirim sebuah variable ke dalam fuction, method atau variable lain sebenarnya yg dikirim adlaah duplikasi valuenya
	// pass by value juga tidak akan mengubah dari dari sumber utamanya/ lebih tepatnya di duplikasi
	// contoh
		// func main(){
		// 	address1 := Address{"subang", "jabar", "indo"}
		// 	address2 := address1

		// 	address2.City = "Bandung"

		// 	fmt.Println(address1) // data subang tidak akan berubah 
		// 	fmt.Println(address2)
		// }

		// type Address struct {
		// 	City, Province, Country string
		// }


// pointer

	// pointer adalah kemampuan membuat referance ke lokasi data di memory yg sama, tanpa menduplikasi data y sudah ada
	// sederhananya dengan kemampuan pointer kita bisa membuat pass by reference
	// pointer tidak akan menduplikasi seperti pass by value

	// cara melakukan pointer di golang
	// operator &
	// utk membuat sebuah variable dengan nilai pinter ke varible yg lain, kita bisa menggunakan operator & diikuti dengan nama variablenya
	// contoh


// Operator * = penanda kalau itu pointer ke data yg dituju
	// saat mengubah variable pointer, maka yg berubah hanya variabel tersebut
	// semua variable yg mengacu ke data yg sama tidak akan berubah
	// jika kita ingin mengubah seluruh variable yg mengacu ke data tsb, kita bisa gunakan operator *


// Function new
	// sebelumnya utk membuat pointer dengan menggunakan operator &
	// go juga memiliki function new yg bisa digunakan utk membuat pointer
	// namun function new hanya mengembalikan poiter ke data kosong artinya tidak ada data di awal