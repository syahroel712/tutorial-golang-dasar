package packageimport

import "fmt"

var Application = "Golang"
var version = "1.0.1"

func SayHello(name string){
	fmt.Println("Hello", name)
}

func sayGoodBye(name string){
	fmt.Println("Good bye", name)
}