package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr2[:]

	arr2[1] = 41
	slice[2] = 69
	fmt.Println(arr2, slice)

	slice = append(slice, 4, 25, 98)
	fmt.Println(slice, arr2)

	s2 := slice[1:]
	s3 := slice[:3]
	s4 := slice[1:3]
	fmt.Println(s2, s3, s4)

	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 69
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	type user struct {
		ID     int
		Role   string
		Gender string
	}

	var u user
	u.ID = 1488
	u.Role = "Engineer"
	u.Gender = "LGBT"
	fmt.Println(u)

	u2 := user{
		ID:     12,
		Role:   "CTO",
		Gender: "Male",
	}
	fmt.Println(u2)
}
