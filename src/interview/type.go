package main

import (
	"fmt"
)

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Println(msg.Name)
	}
}

func main() {

	s := student{
		Name: "name",
	}
	zhoujielun(&s)
}
