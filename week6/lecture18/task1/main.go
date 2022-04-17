package main

import (
	"fmt"
	"time"
	"sync"
)


func main() {

	clock := time.Now()

	fmt.Println(goPrimesAndSleep(100, time.Second))
	fmt.Println(time.Since(clock))

}


func goPrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	var wg sync.WaitGroup
	var m sync.Mutex

	for k := 2; k < n; k++ {

		wg.Add(1)
		go func(k int) {
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i{
						m.Lock()
						res = append(res, k)
						defer m.Unlock()
					}
					wg.Done()
				}
			}
		}(k)
	
	}
	wg.Wait()

	return res
}