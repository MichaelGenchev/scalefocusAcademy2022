package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit)

	tick := time.NewTicker(clearInterval)

	go func() {
		for range tick.C {
			for i := 0; i < bufferLimit; i++ {
				channel <- data
			}
		}
	}()
	return channel
}

func main() {
	out := generateThrottled("foo", 3, time.Second)
	for f := range out {
		log.Println(f)
	}
}


