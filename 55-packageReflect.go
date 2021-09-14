package main

import (
	"fmt"
	"reflect"
)

func main(){
	sample := Sample{"eko", 1}

	// sampleType := reflect.TypeOf(sample)
	var sampleType reflect.Type = reflect.TypeOf(sample)
	// sampleType.Field()

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(1))
	
	// sample.Name = ""
	fmt.Println(IsValid(sample))

}

type Sample struct {
	// contoh biasa
	// Name string
	// Age int

	// contoh pakai structTag
	Name string `required:"true" max:100`
	Age int `required:"true" max:3`

}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == ""{
				return false
			}
		}
	}
	return true
}






// catatan 
// dalam bahasa pemograman, biasanya ada fitur reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi kita sedang berjalan
// hal ini bisa dilakukan dengan package reflect
// reflect sangat berguna ketika kita ingin membuat library yg general shg mudah digunakan
