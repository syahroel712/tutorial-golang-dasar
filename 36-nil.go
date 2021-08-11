package main 

import "fmt"

func main() {

	// cara satu
	// var person map[string]string = nil
	// fmt.Println(person)

	// cara 2
	person := NewMap("")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name" : name,
		}
	}

}


// catatan

// biasanya di dalam bahas lain, object yg belum diinisilisasikan maka secara otomatis nilainya null/nil
// berbeda dengan go, di go saat buat variable dengan tipe data tertentu maka secara otomatis akan dibuatkan default valuenya
// namun di go ada data nil, yaitu kosong
// nil sendiri hanya bisa digunkan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel

// contoh
	// func NewMap(name string) map[string]string{
	// 	if name == "" {
	// 		return nil
	// 	} else {
	// 		return map[string]string{
	// 			"name" : name,
	// 		}
	// 	}
	// }

// contoh
	// var person map[string]string = nil
