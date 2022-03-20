package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)


type TimeSlice []time.Time


func main() {
	
	format := "Jan-02-2006"
	fmt.Println(sortDates(format, "Mar-19-2022", "Dec-03-2021", "Sep-14-2008","Mar-18-2025"))
}


func sortDates(format string, dates ...string) ([]time.Time, error) {
	datesSlice := []time.Time{}

	for _, date := range dates{
		parsedDate, err := time.Parse(format, date)
		if err != nil {
			log.Fatal(err)
		}
		datesSlice = append(datesSlice, parsedDate)
	}
	sort.Sort(TimeSlice(datesSlice))

	return datesSlice, nil
}


func (p TimeSlice) Len() int {
    return len(p) }

// Define compare
func (p TimeSlice) Less(i, j int) bool {
    return p[i].Before(p[j]) }

// Define swap over an array
func (p TimeSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i] }