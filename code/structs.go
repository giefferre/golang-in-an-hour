package main

import (
	"errors"
	"fmt"
)

// START1 OMIT
type Person struct {
	Name    string
	Surname string
	Age     int
}

// STOP1 OMIT

// START2 OMIT
func (p *Person) PrintName() {
	fmt.Println("Hello, my name is:", p.Name)
}

// STOP2 OMIT

// START3 OMIT
func NewPerson(name, surname string, age int) (*Person, error) {
	if (name == "") || (surname == "") {
		return nil, errors.New("Invalid name or surname")
	}

	if age < 0 {
		return nil, errors.New("Invalid age!")
	}

	p := new(Person)
	p.Name = name
	p.Surname = surname
	p.Age = age

	return p, nil
}

// STOP3 OMIT

// START4 OMIT
type Student struct {
	Person
	Class string
}

// STOP4 OMIT

// START5 OMIT
func (s *Student) PrintClass() {
	fmt.Println("My class is:", s.Class)
}

// STOP5 OMIT

// START6 OMIT
func NewStudent(name, surname string, age int, class string) (*Student, error) {
	if (name == "") || (surname == "") || (class == "") {
		return nil, errors.New("Invalid name, surname or class")
	}

	if (age < 15) || (age > 100) {
		return nil, errors.New("Whoa! This guy is too young or old!")
	}

	s := new(Student)
	s.Name = name
	s.Surname = surname
	s.Age = age
	s.Class = class

	return s, nil
}

// STOP6 OMIT

func main() {
	// START7 OMIT
	gianfranco, err := NewPerson("Gianfranco", "Reppucci", 33)
	if err != nil {
		fmt.Println(err)
		return
	}
	// STOP7 OMIT

	if gianfranco.Name != "Gianfranco" {
		// START8 OMIT
		// code using gianfranco...
		gianfranco.PrintName()
		// STOP8 OMIT
	}

	// START9 OMIT
	s, err := NewStudent("John", "Doe", 25, "Computer Science")
	if err != nil {
		fmt.Println(err)
		return
	}

	s.PrintName()
	s.PrintClass()
	// STOP9 OMIT

}
