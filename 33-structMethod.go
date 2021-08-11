package main

import "fmt"

func main(){

	syahrul := Customer{
		Name : "Syahrul",
		Address : "Padang",
		Age : 20,
	}

	syahrul.sayHello("Andi")
	syahrul.sayAge()

}

// struck
type Customer struct {
	Name, Address string
	Age int
}

// struct method
func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name,"my name is", customer.Name)
}

// struct method
func (customer Customer) sayAge(){
	fmt.Println("I am ", customer.Age, "years old")
}


// catatan struct method

// struct adalah tipe data seperti tipe data lainnya(int, bool, string), dia bisa gigunakan sebgai paramter unutk function
// anmun jika kita ingin menambahkan method ke dalam struct, sehingga seakan akan sebuah struct memiliki function
// method adalah function

	// contoh
		// func (customer Customer) sayHello() {
		// 	// fmt.Println("Hello, my name is ", customer.Name)
		// }
		// func (customer Customer) sayAge(){
		// 	// fmt.Println(customer.Age, "years old")
		// }
		// func main(){
		// 	rully := Customer{Name: "Rully", Age: 19}
		// 	rully.sayHello()
		// 	rully.sayAge()
		// }