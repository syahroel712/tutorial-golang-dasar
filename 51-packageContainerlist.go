package main

import (
	"fmt"
	"container/list"
)

func main(){
	
	data := list.New()
	data.PushBack("Syahrul")
	data.PushBack("712")
	data.PushBack("Swimoc")
	data.PushFront("Swimoc")

	// dari depan
	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}

	fmt.Println("----")

	// dari belakang
	for e := data.Back(); e != nil; e = e.Prev(){
		fmt.Println(e.Value)
	}

}



// catatam
// package container/list adalah implementasi struktur data double linked list di go
// data yg saling berkaitan
// hampir mirip array tapi ia saling berkaitan dan ia tidak punya nilai awal dan akhir
// nilai awal dan akhirnya biasanya nil
// cara akses data di linked list harus di loop/iterasi tidak bisa pakai index seperti di array

// cara akses data paling awal = data.front();
// cara akses data setelah paling awal = data.front().Next();
// cara akses data 2 setelah paling awal = data.front().Next().Next();

// contoh
	// func main(){
	// 	data := list.New()
	// 	data.PushBack("Syahrul")
	// 	data.PushBack("712")
	// 	data.PushBack("Swimoc")

	// 	for e := data.Front(); e != nil; e = e.Next(){
	// 		fmt.Println(e.Value)
	// 	}
	// }

