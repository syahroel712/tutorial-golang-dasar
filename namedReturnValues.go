package main

import "fmt"

func main(){

	firstName, middleName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}


func getFullName() (firstName string, middleName string, lastName string){

	firstName = "Syahrul"
	middleName = "712"
	lastName = "."

	return
}










// catatan

// biasanya saat memberi tahu bahwa sebuah function mengembalikan value, maka kita hnay mendeklarasikan tipe data return value di function
// namun kita juga bisa membuat varialbe secara langsung di tipe data return functionnya