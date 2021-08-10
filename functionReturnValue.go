package main

import "fmt"

func main(){

	result := getHello("Syahrul")
	fmt.Println(result)

}

func getHello(name string) string {

	if name == ""{
		return "Hello World" 
	}else{
		return "Hello " + name
	}

}


// catatan

// function bisa mengembalikan data
// utk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tsb
// jika function tsb kita deklarasikan dengan tipe data pengembalaian, maka wajib di dalam functionnya kita harus mengembalikan data
// utk mengembalikan data dari function, kita bisa gunakan kata kunci return diikuti dengan datanya
	// func getHello(name string) string {
	// 	return "Hello" + name
	// }



