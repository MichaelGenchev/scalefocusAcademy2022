package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer  int
	timeout time.Duration
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	bc := &BufferedContext{ctx, bufferSize, timeout}

	return bc

}
func (bc *BufferedContext) Done() <-chan struct{} {
	ctx, cancel := context.WithTimeout(bc, bc.timeout)
	defer cancel()
	return ctx.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	ch := make(chan string, bc.buffer)
	for {
		if len(ch) != cap(ch) {
			ch <- "bar"
			time.Sleep(time.Millisecond * 60)
		}
		fn(bc.Context, ch)
		close(ch)
		return
	}
}

func main() {
	ctx := NewBufferedContext(5*time.Second, 100)
	fmt.Println("start")
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200)
				fmt.Println("bar")
			}
		}
	})
}