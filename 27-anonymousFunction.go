package main

import "fmt"

func main(){

	// cara 1, dengan mendekrasikan ke variable
	blacklist := func(name string) bool {
		return name == "Admin"
	}

	registerUser("Admin", blacklist)
	registerUser("Syahrul", blacklist)

	// cara 2 secara langsung anonymous function di parameter
	registerUser("root", func(name string) bool{
		return name == "root"
	})
	registerUser("Syahrul", func(name string) bool{
		return name == "root"
	})


}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}








// catatan anonymous function
// sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tsb
// namun kadang , lebih mudah membuat function secara langusng di variable atau paramter tanpa harus membuat function terlebih dahulu
// hal tersebut dinamakan anonymous function , / function tanpa nama