package main

import "fmt"

func main(){

	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)



	// // append code
	// var slice2 = months[10:]
	// fmt.Println(slice2)
	
	// var slice3 = append(slice2, "Bulan Baru")
	// fmt.Println(slice3)
	// slice3[1] = "Bukan Desember"
	// fmt.Println(slice3)
	// fmt.Println(slice2)
	// fmt.Println(months)


	// make slice code
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Syahrul"
	newSlice[1] = "712"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))


	// copy slice code
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// hari buat array dan slice
	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}






// catatan
// tipe data slice adalah potongan dari data array
// slice mirip denan array yg membedakan adalah ukuran slice bisa berubah
// slice dan array selalu terkoneksi, dimana slice dalah data yg mengakses sebagian atau seluruh data di array
// saat mengubah isi array di slice maka data di array juga berubah, atau sebaliknya

// detail tipe data slice
	// tipe data sliice memiliki t3 data, pointer, length, dan capacity
	// pointer adalah penujuk data pertama di array pada slice
	// length adalah panjang dari slice
	// capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

// membuat slice dari array
	// array[low:high] 	=> membuat slice dari data array dimulai index low sampai index sebelum high
	// array[low:] 		=> membuat slice dari array dimulai index low sampai index akhir di array
	// array[:high] 	=> membuat slide dari array dimulai index 0 sampa index seblum high
	// array[:]			=> membuat slice dari array simulai dari index 0 sampai index akhir di array

// function slice
	// len(slice)							=> mendapatkan panjang array
	// cap(slice) 							=> mendapatkan kapasitas array
	// append(slice, data) 					=> membuat slice baru dengan menambah daata ke posisi terakhir slice, jika kapasitas penuh maka 											akan membuat array baru
	// make([]TypeData,length, capacity) 	=> membuat slice baru
	// copy(destination, source)			=> menyalin slice dari source ke destination