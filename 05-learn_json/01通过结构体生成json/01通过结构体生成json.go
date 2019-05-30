package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写，才能生成json,但是json的键是大写的
// type IT struct {
// 	Name    string
// 	Age     int
// 	City    string
// 	Country string
// }
//修改json的键，是小写的
type IT struct {
	Name    string `json:"-"`       //此字段不会输出到屏幕
	Age     int    `json:",string"` //将整型转换为string
	City    string `json:"city"`
	Country string `json:"country"`
}

/*
buf= {
 "Age": "25",
 "city": "Guangzhou",
 "country": "China"
}
*/
func main() {
	s := IT{"Paul", 25, "Guangzhou", "China"}

	// buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("出错")
		return
	}
	fmt.Println("buf=", string(buf))

}
