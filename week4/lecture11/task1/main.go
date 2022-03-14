package main

import (
	// "fmt"
	// "time"
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, ch chan int, num int){
	defer wg.Done()
	ch <- num
}

func monitorWorker(wg *sync.WaitGroup, ch chan int){
	wg.Wait()
	close(ch)
}

func processEven(inputs []int) chan int {
	wg := &sync.WaitGroup{}
	evenCh := make(chan int)
	
	for i := 0; i < len(inputs); i++ {
		num := inputs[i]
		if num % 2 == 0 {
			wg.Add(1)
			go worker(wg, evenCh, num)
		}
	}
	go monitorWorker(wg, evenCh)

	
	return evenCh
}

func processOdd(inputs []int) chan int {
	wg := &sync.WaitGroup{}
	oddCh := make(chan int)
	
	for i := 0; i < len(inputs); i++ {
		num := inputs[i]
		if num % 2 != 0 {
			wg.Add(1)
			go worker(wg, oddCh, num)
		}
	}
	go monitorWorker(wg, oddCh)

	return oddCh
}

func main() {
	
	inputs := []int{1, 17, 34, 56, 2, 8}
	
	// evenCh := processEven(inputs)

	// for i := range evenCh {
	// 	fmt.Println(i) 
	// }

	oddCh := processOdd(inputs)

	for i := range oddCh{
		fmt.Println(i)
	}


}
