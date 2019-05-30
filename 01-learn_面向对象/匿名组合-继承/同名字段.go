package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
	name string //和Person有同名的变量了
}

func main() {
	var s Student
	//默认规则（就近原则）
	s.name = "mike" //操作的是Student的name,不是Person的name
	s.sex = 'm'
	s.age = 18
	s.addr = "obj"
	s.Person.name = "Paul"
	fmt.Printf("s=%+v\n", s)
}
