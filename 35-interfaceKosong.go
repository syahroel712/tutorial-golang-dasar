package main

import "fmt"

func main(){

	var data interface{} = Ups(3)
	fmt.Println(data)

}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}



// catatan
// go bukanlah bahasa yg berorientasi object
// biasanya dalam pemogramman oop, ada satu data parent di puncak yg dianggap sebagai semua implementasi data yg ada d bhahas pemogramana tsb
// contoh di java ada java.lang.Object
// utuk menangani kasus seperti ini, di go kita bisa menggunakan interface kosong
// interface kosong adalah interface yg tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

// penggunaan interface kosong di golang
// ada banay contoh penggunaan interface kosong, 
	// fmt.Println(a....inteface{})
	// panic(v....inteface{})
	// recover() v....inteface{}
	// dll

// contoh
	// func Ups() inteface{}{
	// 	// return 1
	// 	// return true
	// 	return "Ups"
	// }

	// func main(){
	// 	kosong := Ups()
	// 	fmt.Println(kosong)
	// }


// interface kosogn biasa digunakan jika kita ingin menerima semua type data di interface tersebut