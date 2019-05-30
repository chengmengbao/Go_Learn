package main

import (
	"encoding/json"
	"fmt"
)

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
	//创建一个map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// fmt.Println("m=", m)
	// fmt.Printf("m=%+v\n", m)

	// var s string
	// s = m["company"].(string)
	// fmt.Println("s= ", s)
	var str string
	//类型断言
	for key, value := range m {
		// fmt.Printf("%v===>%v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string，内容为%s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool，内容为%v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64，内容为%v\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]stiring，内容为%v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]stiring，内容为%v\n", key, data)
		}
	}
}
