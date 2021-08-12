package main

import "fmt"

func main(){

	andi  := Man{"Andi"}
	andi.Married()
	fmt.Println("Di main akan diubah sesuai yg di method",andi.Name)

}

type Man struct {
	Name string
}

// tambahanan *(pointer) pada func
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("di method akan di ubah jadi", man.Name)

}


// catatan
// walaupun method akan menempel di struct, tapi sebenarnya data struct yg diakses di dalam method adalah pass by value
// sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
