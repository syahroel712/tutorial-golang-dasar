package main

import "fmt"

func main(){

	var names [3]string

	names[0] = "Syahrul"
	names[1] = "712"
	names[2] = "."

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	
	
	var angka = [3]int{
		20,
		21,
		22,
	}
	fmt.Println(angka)
	fmt.Println(len(angka))

}


// catatan

// array adalah tipe data yg berisikan kumpulan dengna tipe data yg sama
// saat membuat array, kita perlu menentukan jumlah data yg bisa ditamping oleh array
// daya tampung array tidak bisa bertambah setealh dibuat
// cara mengambil isi array adalah degan menggunakan index, dimulai dari 0


// function array
	// len(array) => utk mencari panjang array
	// array[index] => mendapatkan data di posisi index
	// array[index] = value => mengubah data di posisi index

