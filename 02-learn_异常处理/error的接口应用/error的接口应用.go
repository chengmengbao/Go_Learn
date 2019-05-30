package main

import "fmt"

// import "errors"

// func MyDiv(x, y int) (res int, err error) {
// 	if y == 0 {
// 		err = errors.New("分母不能为0")
// 	} else {
// 		res = x / y
// 	}
// 	return
// }

// func main() {
// 	res, err := MyDiv(10, 4)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 	} else {
// 		fmt.Println("res:", res)
// 	}
// }
type I interface {
	M(name string) //接口的方法声明
}

type T struct {
	name string
}

func (t *T) M(name string) {
	// fmt.Printf("Hi, my name is %s\n", t.name)
	t.name = name
}

func main() {
	p := &T{"mike"}
	var i I = p
	// f := I.M
	// f(i, "Paul")
	i.M("xxx")
	fmt.Println(i.(*T).name)
	// t := &T{"foo"}
	// f := (*T).SayHi
	// f(t) // Hi, my name is foo
}
