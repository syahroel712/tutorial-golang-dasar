package main 

import "fmt"

func main(){

	var name string
	name = "Syahrul"
	fmt.Println(name)
	
	name = "Syahrul 712"
	fmt.Println(name)

	age := 20
	fmt.Println(age)

	var (
		firstName = "Syahrul"
		lastName = "."
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}


// catatan
// variable tempat menyimpan data
// variable digunakan agar bisa akses data yg sama dimanapun dibutuhkan
// di go variable hanya bisa menyimpan tipe data yg sama, jika berbeda maka kita harus buat variable baru
// untuk buat variable gunkan kata kunci var lalu nama varibale dan tipe datanya
	// var angka int8
	// var name string

// jika buat variabel maka harus tulis tipe data, jika langsung mengisi nilai variable maka tidak perlu menuliskan tipe datanya
	// var name = "syahrul"
	// var age = 20

// kata kunci var tidak wajib jika kita langusung mmengisi nilai dari varibel tersebut dengan catatan kita menggunakan kata kunci :=  saat menginisialisasikan data variable tersebut
	// age := 20
	// name := "syahrul"


// di go ktia bisa buat variable secara sekaligus banyak
// code akan lebih dimudaj=h dibaca
	// var (
	// 	firstName = "Syahrul"
	// 	lastName = "."
	// )