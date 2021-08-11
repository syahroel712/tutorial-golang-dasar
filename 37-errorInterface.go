package main

import (
	"errors"
	"fmt"
)

func main(){
	// contoh
	// var contohError error = errors.New("Test Error")
	// fmt.Println(contohError)


	hasil, err := Pembagian(100, 0)
	if err == nil{
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}

func Pembagian(nilai int, pembagi int) (int, error){
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	}else{
		result := nilai / pembagi
		return result, nil
	}
}



// catatan
// golang memilki interface yg digunakan sebgai kontrak untk membuat error, nama interface tsb adalah error

// ex 
	// type error interface {
	// 	Error() string
	// }

// membuat error
// utk membuat errro, kita tak perlu manual
// golang sudah menyediakan library utk membuat helper secara mudah, yg terdapat di package errors
