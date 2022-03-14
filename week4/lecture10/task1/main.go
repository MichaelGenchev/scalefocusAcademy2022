package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex

	text string
}
func incrementFoo(cp *ConcurrentPrinter, times int){
	for i := 0; i < times/2; i++{
		cp.Lock()
		cp.text += "foo"
		cp.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
	cp.Done()
}

func incrementBar(cp *ConcurrentPrinter, times int){
	for i := 0; i < times/2; i++{
		cp.Lock()
		cp.text += "bar"
		cp.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
	cp.Done()

}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.Add(1)
	go incrementFoo(cp, times)
}
func (cp *ConcurrentPrinter) printBar(times int) {
	cp.Add(1)
	go incrementBar(cp, times)
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.printFoo(times)
	cp.printBar(times)
	cp.Wait()
	fmt.Println(cp.text)

}


