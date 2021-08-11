package main

import "fmt"

func main(){

	person := Person{Name : "Andi"}
	sayHello(person)


	cat := Animal{Name: "Oyen"}
	sayHello(cat)

}

// interface
type HasName interface{
	GetName() string
}
// 
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// struct
type Person struct {
	Name string
}
// struck method, GetName otomatis jadi milik interface
func (person Person) GetName() string {
	return person.Name
}


// struct ke 2
type Animal struct {
	Name string
}
func (animal Animal) GetName() string {
	return animal.Name
}





// catatan 
// interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
// sebuah interface berisikan definisi2 method
// biasanya interface digunakan sebagai kontrak
	// contoh
	// type HasName interface {
	// 	GetName() string
	// }

	// func sayHello(hasName HasName) {
	// 	fmt.Println("Hello", hasName.GetName())
	// }


// implementasi interface
// setiap tipe data yg sesuai dengan kontrak interface, secar otomatis dianggap sebagai interface tsb
// sehingga kita tidak perlu mengimplementasikan interface secara manual
// hal ini berbeda dengan bahasa lain ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana
	
	// contoh: pada struct Person ada sebuah struct method yg namanya GetName, maka GetName tersebut secara otomatis masuk ke interface dari HasName pada contoh diatas

	// type Person struct {
	// 	Name string
	// }
	// func (person Person) GetName() string {
	// 	return person.Name
	// }


