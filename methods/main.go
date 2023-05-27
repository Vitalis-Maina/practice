package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name string
	age  int
	Courses
}

var choice int
var list = make([]Student, choice)

type Courses struct {
	courseName string
	duration   string
}

func ListStudents() []Student {

	fmt.Println("ALL STUDENTS IN THE DATABASE")

	return list
}

func (s *Student) AddStudent() {

	fmt.Println("Enter numeber you want to add: ")
	fmt.Scan(&choice)

	fmt.Println("Enter Student details: ")
	var (
		Name       string
		age        int
		courseName string
		duration   string
	)
	var listOfStudents []Student

	for i := 1; i <= choice; i++ {
		fmt.Scan(&Name, &age, &courseName, &duration)

		list = []Student{
			{
				Name: Name,
				age:  age,
				Courses: Courses{
					courseName: courseName,
					duration:   duration,
				},
			},
		}
		listOfStudents = append(listOfStudents, list...)
		list = listOfStudents
	}
	fmt.Println("STUDENT ADDED SUCCESSFULLY")

}

func (s Student) DeleteStudent(name string) {
	indexToDelete := -1

	for i, item := range list {
		if item.Name == name {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		list = append(list[:indexToDelete], list[indexToDelete+1:]...)
	}

	fmt.Println("STUDENT DELETED SUCCESSFULLY")
}

func (s *Student) Modify() {
	fmt.Println("modify")
	fmt.Println("Enter name of student:")
	var Modname string
	fmt.Scan(&Modname)

	for i := range list {
		if list[i].Name == Modname {
			var (
				name           string
				age            int
				courseName     string
				courseDuration string
			)
			fmt.Scan(&name, &age, &courseName, &courseDuration)

			list[i].Name = name
			list[i].age = age
			list[i].courseName = courseName
			list[i].duration = courseDuration

		}

	}
	fmt.Println("MODIFICATION SUCCESSFULL")

}
func main() {

	var students Student
	fmt.Println("WELCOME TO OUR PORTAL")
	for {
		fmt.Printf("\t\npick \n1: ADD  STUDENT \n2: SHOW STUDENTS \n3: DELETE STUDENT \n4: MODIFY STUDENT \n6: EXIT\n")
		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:
			students.AddStudent()

		case 2:
			fmt.Println(ListStudents())

		case 3:
			fmt.Println("Enter name to delete")
			var name string
			fmt.Scan(&name)
			students.DeleteStudent(name)

		case 4:
			students.Modify()

		case 0:
			os.Exit(0)

		}
	}

}
