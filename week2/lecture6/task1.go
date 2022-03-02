package main

import "fmt"


type MagicList struct {
	LastItem *Item
}
type Item struct {
	Value int
	PrevItem *Item
}

func add(l *MagicList, value int) {
	item := Item{Value: value}
	if l.LastItem == nil {
		l.LastItem = &item
	}else {
		item.PrevItem = l.LastItem
		l.LastItem = &item
	}
}

func toList(l *MagicList) []int {
	var current = l.LastItem
	var slice = make([]int, 0)

	slice = append(slice, current.Value)
	for true {
		if current.PrevItem == nil {
			break
		} else {
			current = current.PrevItem
			fmt.Println(current.Value)
			slice = append(slice, current.Value)
		}

	}

	return reverse(slice)
}
func reverse(input []int) []int {
	inputLen := len(input)
	output := make([]int, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}
func main() {
	l := &MagicList{}
	add(l, 1)
	add(l, 2)
	add(l, 3)
	add(l, 4)
	fmt.Println(l.LastItem.Value)
	result := toList(l)
	fmt.Println(result)
}