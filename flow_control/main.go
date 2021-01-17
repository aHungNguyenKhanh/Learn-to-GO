package main

type User struct {
	ID        int
	FirstName string
	Lastname  string
}

type HTTPRequest struct {
	Method string
}

func main() {
	for i := 0; i < 5; i++ {
		println(i)
		if i == 3 {
			continue
		}
		println("continuing")
	}

	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	ports := map[string]int{"http": 80, "ssh": 22}
	for k, v := range ports {
		println(k, v)
	}

	u1 := User{
		ID:        1,
		FirstName: "Hung",
		Lastname:  "Nguyen",
	}
	u2 := User{
		ID:        2,
		FirstName: "Phuc",
		Lastname:  "Doan",
	}
	if u1.ID == u2.ID {
		println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user.")
	} else {
		println("Different user.")
	}

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("GET request")
		fallthrough
	case "POST":
		println("POST request")
	case "DELETE":
		println("DELETE request")
	case "PUT":
		println("PUT request")
	default:
		println("Undefined method")
	}
}
