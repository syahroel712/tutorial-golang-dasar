package main

import "fmt"

func main(){

	// contoh 1
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)


	// contoh dengan switch
	switch value := result.(type){
		case string :
			fmt.Println("Value", value, "is string")
		case int :
			fmt.Println("Value", value, "is int")
		case bool :
			fmt.Println("Value", value, "is bool")
		default : 
			fmt.Println("Unknown Type")
	}

}

func random() interface{}{
	return "ASd"
}

// catatan 
// type assertions merupakan kemampuan merubah tipe data menjadi tipe data yg diinginkan
// fitur ini sering digunakn ketika bertemu dengan data interface kosong
// pastikan saat melakukan type assertions data yg di return dengan yg di buat sama
// contoh
	// func random() interface{}{
	// 	return "Ok"
	// }

	// func main() {
	// 	result := random()
	// 	resultString := result.(string) // proses type assertions
	// 	fmt.Println(resultString)
		
	// 	resultInt := result.(int) // panic, karena hasil datanya adlah string
	// 	fmt.Println(resultInt)
	// }



// type assertions menggunakan switch
// saat salah menggunakan type assertion, maka bisa berakibat terjadi panic di aplikasi 
// jika panic dan tidak ter- recover, maka otomatis program kita akan berhenti
// agar lebih aman , sebiknya gunakan switch expression utk melakukan type assertions

// contoh
	// func main(){
	// 	result := random()
	// 	switch value := result.(type) {
	// 		case string :
	// 			fmt.Println("String", value)
	// 		case int :
	// 			fmt.Println("Int", value)
	// 		default:
	// 			fmt.Println("Unknown")
	// 	}
	// }
