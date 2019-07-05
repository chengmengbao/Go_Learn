package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`  //二次编码，此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

//经典错误示例
//type User struct {
//	Name string `json:"name"`
//}

func main() {
	jsonBuf := `
	{
		"company": "itcast",
		"subjects": [
			"Go",
			"C++",
			"Python",
			"Test"
		],
		"isok": true,
		"price": 666.666
	}
	`
	var tmp IT
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// fmt.Println("tmp=", tmp)
	fmt.Printf("tmp=%+v\n", tmp)
	//tmp={Company:itcast Subjects:[Go C++ Python Test] IsOk:true Price:666.666}

	//经典错误示例
	//var u User
	//json.Unmarshal([]byte(`{"name":"polaris"}`), &u)
	//fmt.Printf("u.Name=%+v\n", u.Name)
}

//经典错误解释
//1.
//var u *User
//json.Unmarshal([]byte(`{"name":"polaris"}`), u)
//fmt.Printf("u.Name=%+v\n", u.Name)
//以上会直接报错：panic: runtime error: invalid memory address or nil pointer dereference
//2.
//var u *User
//json.Unmarshal([]byte(`{"name":"polaris"}`), &u)
//fmt.Printf("u.Name=%+v\n", u.Name)
//正常打印出u.Name=polaris
//问题：var u *User， u不就是指针吗，为什么在json.Unmarshal(`{"name":"polaris"}`), &u)里面，还要&u，相当于取指针的地址？请教星主
//3.
//var u User
//json.Unmarshal([]byte(`{"name":"polaris"}`), &u)
//fmt.Printf("u.Name=%+v\n", u.Name)
//正常打印出u.Name=polaris
