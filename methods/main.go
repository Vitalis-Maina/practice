package main

import "fmt"

type Student struct {
	Name string
	age  int
}

type Courses struct {
	Bsc  string
	Bbit string
}

// there will be no methods overload since the
// receivers are of different types
func (s Student) Info() {

	fmt.Println("name is ", s.Name, "age is", s.age)
}

func (c Courses) Info() {
	fmt.Println("courses offered", c.Bsc, c.Bsc)
}
func main() {

}
