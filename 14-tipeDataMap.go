package main 

import "fmt"

func main(){

	person := map[string]string{
		"name" : "Syahrul",
		"address" : "Padang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	
	person["title"] = "Programmer"
	fmt.Println(person)


	// function map
	var book map[string]string = make(map[string]string)
	book["title"] = "Learn GO"
	book["author"] = "Syahrul"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))



}



// catatan

// pada array/slice , untuk akses data kita gunkana index number dimulai dari 0
// map adalah tipe data lain yg berisikan kumpulan data yg sama, namun kita bisa menentukan jenis tipe data yg akan kita gunakan
// sederhanananya, map adalah tipe data kumpulan key-value(kata kunci - nilai)dimana kata kuncinya hasur unik, tidak boleh sama
// /berbeda dengan aray/slice, jumlah data yg kita masukkkan ke dalam map boleh sebanyaknya, asalkan kata kuncinya beda, jika sama maka otomatis data yg sebelumnya akan di ganti dengan data yg baru

// function map
	// len(map)  					=> untuk mendapatkan jumlah data di map
	// map[key]						=> mengambil data di map dengan key
	// map[key] = value 			=> mengubah data di map dengan key
	// make(map[TypeKey]TypeValue) 	=> membuat map baru
	// delete(map, key)				=> menghapus data di map dengan key 