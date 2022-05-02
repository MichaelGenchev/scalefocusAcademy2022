package main


import "testing"


func TestGrouping(t *testing.T) {
	results := GroupBy([]Order{    {Customer: "John", Amount: 1000},    {Customer: "Sara", Amount: 2000},    {Customer: "Sara", Amount: 1800},    {Customer: "John", Amount: 1200},}, func(o Order) string { return o.Customer } )
	
	want := map[string][]int{"John": []int{1000,1200},"Sara": []int{2000, 1800}}

	if results == nil {
		t.Errorf("GroupBy function returned %v, want %v", results, want)
	}
	
}