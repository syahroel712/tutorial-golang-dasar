package main

import "fmt"

func main(){
	sayGoodbye := getGoodBye

	result := sayGoodbye("Syahrul")

	fmt.Println(result)

}


func getGoodBye(name string) string {
	return "Good bye " + name
}


// catatan

// function adalah first class citizen di golang	
//  function juga merupkan tipe data, dan bisa disimpan di dalam variable	
	// contohnya di function main di variable goodbye kita menjadikan Func getGoodBye sebagai value 
	// func getGoodBye(name string) string {
		// return "Good Bye" + name
	// }	

	// func main(){
		// goodbye := getGoodBye
	// 	fmt.Println(goodbye("Syahrul"))
	// }



