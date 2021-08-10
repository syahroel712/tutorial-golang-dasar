package main

import "fmt"

func main(){
	// defer
	// runApp()

	// panic
	// runApp(true)

	// recover
	runApp(true)
	fmt.Println("Syahrul")
}

// // defer
// func runApp(){
// 	// defer letaknya diawal
// 	defer logging()
// 	fmt.Println("Run app")
// }
// func logging(){
// 	fmt.Println("Selesai memanggil function")
// }


// // panic
// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("Aplikasi error")
// 	}
// 	fmt.Println("Aplikasi berjalan")
// }
// func endApp(){
// 	fmt.Println("Aplikasi selesai")
// }

// // recover
func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
}
func endApp(){
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error", message)
	}
	fmt.Println("Aplikasi selesai")
}





// cattaan

// Defer
// defer function adalah function yg bisa kita jadwalkan utk dieksekusi setelah sebuah function selesai di eksekusi
// defer function akan selalu di eksekusi walaupun terjadi error di function yg di eksekusi
	// contoh
		// func logging(){
			// fmt.Println("Selesai memanggil function")
		// }

		// func runApp(){
		// 	defer logging()
		// 	fmt.Println("Run App")
		// }

		// func main(){
		// 	runApp()
		// }



// panic
// panic adalah function yg bisa digunakan utk menghentikan program
// panic biasanya dipanggil ketika terjadi error pada sat program berjalan
// saat panic dipanggil, program akan berhenti namun defer function tetap dieksekusi
	// contoh
		// func endApp(){
		// 	fmt.Println("End App")
		// }

		// func runApp(error bool){
		// 	defer endApp()
		// 	if error {
		// 		panic("ERROR")
		// 	}
		// }




// Recover
// reccover adlah function yg kita gunakan utk menangkap data panic
// dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan / dengan recover, panic tidak akan menghentikan aplikasi
	// contoh
		// func runApp(error bool){
			// 	defer endApp()
			// 	if error{
			// 		panic("Error")
			// 	}
		// }

		// func endApp(){
			// 	message := recover()
			// 	if message != nil {
			// 		fmt.Println("Terjadi error", message)
			// 	}
			// 	fmt.Println("Aplikasi selesai")
		// }


// recover sebaiknya di letakkan dalam defer function 


