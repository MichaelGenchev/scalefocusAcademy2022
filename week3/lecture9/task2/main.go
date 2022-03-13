package main

import "fmt"


// type Action func() error

// func SafeExec(a Action) error {
// 	err := a()
// 	if err != nil {
// 		return fmt.Errorf("safe exec: %w", err)
// 	}
// 	return nil
// }

func f() (err error) {
	defer func() {
	if r := recover(); r != nil {
	err = fmt.Errorf("Recovered in f: %s", r)
	}
	}()
	panic("Boom!")
}


func main() {
	f()

}