package main

import (
	"github.com/amehrotra/interface/student"
	"github.com/amehrotra/interface/action"
)

func main()  {
	s := student.New(1,21,"Aryan",5)
	action.Perform(s)
}