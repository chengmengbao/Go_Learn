package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

type Cat struct {
	*Animal
	Type string
}

func TestInherit(a *Cat) {
	fmt.Println(a.Name)
}

func main() {
	cat := &Cat{
		Animal: &Animal{Name: "xiaohua", Age: 12},
		Type:   "波斯猫",
	}
	TestInherit(cat)
}

//func TestInherit01(a *Animal) {
//	fmt.Println(a.Name)
//}
//
//func main01() {
//	cat := &Cat{
//		Animal: &Animal{Name: "xiaohua", Age: 12},
//		Type:   "波斯猫",
//	}
//	TestInherit01(cat) //报错：cannot use cat (type *Cat) as type *Animal in argument to TestInherit01
//}
