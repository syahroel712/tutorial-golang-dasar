package main

import (
	"fmt"
	"sort"
)

func main(){
	users := []User {
		{"Lena", 18},
		{"Lani", 27},
		{"Lina", 16},
		{"Lala", 25},
		{"Lili", 28},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (value UserSlice) Len() int{
	return len(value)
}

func (value UserSlice) Less(i, j int) bool{
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int){
	value[i], value[j] = value[j], value[i]
} 



// catatan
// package yg berisikan utilitas utk proses pengurutan
// agar data bisa diurutkan, kita harus mengimplementasikan kontrakdi interface sort.Interface

