package main

import "fmt"

func main(){

	var alamat Address = Address{
		City : "subang", 
		Province : "jabar", 
		Country: "test",
	}
	// tandai dengan & supaya bisa menjadi pointer
	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)


}

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){ //tambakan *(pointer) utk memberitahu bahawa ini merupakan data pointer
	address.Country = "Indonesia"
}


// catatan pointer di function

// saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
// olehkarana itu, jika kita mengubah data didalm function, data yg aslinya tidak akan berubah
// hal ini membuat variable menjadi aman karena tidak akan bisa di ubha
// namun kadang kita ingin membuat function yg bisa mengubah data asli paramter tesbut
// utk melakukan ini, kita juga bisa menggunakan pointer di function
// utk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
// usahakan jika menggunakan data struct yg besar, dan dijadikan paramter maka gunakanlah pointer supaya penggunaan memori tidak bengkak 