package main

import "fmt"




func isItLeapYear(year int) bool {
	if( (year%400==0 || year%100!=0) &&(year%4==0)){
		return true
	}else {
		return false
	}
	
}
func daysInMonth(month int, year int) (int, bool) {
	isLeapYear := isItLeapYear(year)
	switch month {
		case 1,3,5,7,8,10,12:
			return 31, false
		case 2:
			if isLeapYear {
				return 29, false
			}else {
				return 28, false
			}
		case 4, 6, 9, 11:
			return 30, false
	}
	return 0, true
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	year := 4
	month := 2
	// scanner.Scan()
	// month := scanner.Text()
	result, err  := daysInMonth(month, year)
	if err {
		fmt.Println("Invalid month: ", month)
	}else {
		fmt.Println(result)
	}

	
}