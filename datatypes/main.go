package main

import (
	"fmt"
)

const (
	myDailyWorkHours = iota + 7
	myDailySleepHours
	myAge = 4<<iota + 9
)

const (
	newConst = iota
)

func main() {
	var i int
	i = 69
	fmt.Println(i)

	var f float64 = 1.41
	fmt.Println(f)

	myName := "Hung"
	fmt.Println(myName)

	b := false
	fmt.Println(b)

	c := complex(6, 9)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var firstName *string = new(string)
	*firstName = "Hung"
	fmt.Println(*firstName)

	ptr := &myName
	fmt.Println(ptr, *ptr)

	myName = "Nguyen"
	fmt.Println(ptr, *ptr)

	const pi int = 3
	fmt.Println(float64(pi) + 0.14)

	fmt.Println(myDailyWorkHours, myDailySleepHours, myAge, newConst)
}
