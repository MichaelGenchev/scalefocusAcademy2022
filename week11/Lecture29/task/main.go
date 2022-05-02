package main

import (
	"fmt"
	"reflect"
)
func main() {

	results := GroupBy([]Order{    {Customer: "John", Amount: 1000},    {Customer: "Sara", Amount: 2000},    {Customer: "Sara", Amount: 1800},    {Customer: "John", Amount: 1200},}, func(o Order) string { return o.Customer } )

	fmt.Println(results)
}


type Order struct {    
	Customer string    
	Amount int
}


func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]int {
	result := make(map[U][]int)
	for _, person := range col {
		name := keyFn(person)
		amount := getFieldInteger(person, "Amount")
		result[name] = append(result[name], amount)
	}
	return result
}

func getFieldInteger(e any, field string) int {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}