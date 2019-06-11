package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "Inter"
	m["subjects"] = []string{"Go", "C++", "Python", "Java"}
	m["isok"] = true
	m["price"] = 666.666

	// res, err := json.Marshal(m)
	res, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("res=", string(res))
}

//打印结果
//res= {
//	"company": "Inter",
//	"isok": true,
//	"price": 666.666,
//	"subjects": [
//		"Go",
//		"C++",
//		"Python",
//		"Java"
//	]
//}
