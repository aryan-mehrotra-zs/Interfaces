package student

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
	Gpa  float64
}

func New(id, age int, name string, gpa float64) Student {
	return Student{id, name, age, gpa}
}

func (s Student) Talk() string {
	return fmt.Sprintf("Student Id : %v \nName       : %s", s.Id, s.Name)
}
