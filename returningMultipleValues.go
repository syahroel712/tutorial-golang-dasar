package main

import "fmt"

func main(){
	firstName, _, age := getFullName()

	fmt.Println(firstName)
	// fmt.Println(lastName)
	fmt.Println(age)

}

func getFullName()(string, string, int8){
	return "Syahrul", ".", 20
}











// catatatn

// function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// untuk memberitahu jika function mengembalikan multiple value, kita hasru menuliskan semua tipe data return valuenya di function
	// contoh fucntion
		// func getFullName() (string, string){
			// return "Syahrul", "."
		// }
	// contoh cara panggil
		// func main(){
		// 	firstName, lastName := getFullName()
		// 	fmt.Println(firstName, lastName)
		// }


// menghiraukan return value
// multiple return value wajib ditangkap semua valuenya
// jika kita inggin menghiraukan reutn value tsb, kita bisa gunakan tanda _(underscore)
	// contoh fucntion
		// func getFullName() (string, string){
			// return "Syahrul", "."
		// }
	// contoh cara panggil
		// func main(){
		// 	firstName, _ := getFullName()
		// 	fmt.Println(firstName)
		// }


