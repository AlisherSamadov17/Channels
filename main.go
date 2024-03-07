package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1 -masala

func sendrandN(ch chan <- int)  {
	ch <-rand.Intn(100)
}
func main()  {
	ch :=make(chan int)
    go sendrandN(ch)
	sum := <-ch
	time.Sleep(time.Second)
	fmt.Println(sum)
}


// //__2-masala__
// func randomN(n int, ch chan<- int) {
// 	defer close(ch)
// 	for i := 0; i < n; i++ {
// 		ch <- rand.Intn(10)
// 	}
// }

// func main() {
// 	n := 3
// 	ch := make(chan int)

// 	go randomN(n, ch)

// 	for v := range ch {
// 		fmt.Println(v)
// 	}
