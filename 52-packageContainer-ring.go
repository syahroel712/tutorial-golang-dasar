package main

import (
	"fmt"
	"container/ring"
	"strconv"
)

func main(){
	// data := ring.New(5)
	var data *ring.Ring =  ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

}

// catatan
// adalah implementasi struktur data circular list
// Circular list adalah  struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
// contoh
	// func main() {
	// 	data := ring.New(5)
	// 	for i := 0; i < data.Len(); i++ {
	// 		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
	// 		data = data.Next()
	// 	}

	// 	// melakukaln iterasi degan function bawaan ring yaitu Do
	// 	data.Do(func(value interface{}) {
	// 		fmt.Println(value)
	// 	})
	// }
