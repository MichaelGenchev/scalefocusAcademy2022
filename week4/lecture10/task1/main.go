package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}


func (cp *ConcurrentPrinter) printFoo(times int) {
	go func() {
		for i := 0; i < times; i++ {
			defer cp.Done()
			cp.Add(1)
			if  i % 2== 0 {
				fmt.Print("foo")
			}
			time.Sleep(time.Millisecond)
		}
	}()
}
func (cp *ConcurrentPrinter) printBar(times int) {
	fmt.Print("")
	go func() {
		for i := 0; i < times; i++ {
			defer cp.Done()
			cp.Add(1)
			if  i % 2 != 0 {
				fmt.Print("bar")
			}
			time.Sleep(time.Millisecond)
		}
	}()

}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.printFoo(times)
	cp.printBar(times)
	cp.Wait()
}


