package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlystudents struct {
	XMLName     xml.Name  `xml:"students"`
	Version     string    `xml:"version,attr"`
	Students    []student `xml:"student"`
	Description string    `xml:",innerxml"`
}
type student struct {
	XMLName     xml.Name      `xml:"student"`
	StudentName string        `xml:"studentName"`
	Age         int           `xml:"age"`
	Sex         string        `xml:"sex"`
	Books       Recurlybookss `xml:"books"`
}

type Recurlybookss struct {
	XMLName     xml.Name `xml:"books"`
	Version     string   `xml:"version,attr"`
	Books       []book   `xml:"book"`
	Description string   `xml:",innerxml"`
}

type book struct {
	XMLName  xml.Name `xml:"book"`
	BookName string   `xml:"bookName"`
	Price    string   `xml:"price"`
}

func main() {
	file, err := os.Open("F:\\code\\go\\src\\Go_Learn\\14-learn_XML处理\\02XML处理复杂版\\servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlystudents{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}

//运行结果
/*
{{ students} 1 [{{ student} 韩茹 31 female {{ books}  [{{ book} 红与黑 55.8} {{ book} 呼啸山庄 99}]
            <book>
                <bookName>红与黑</bookName>
                <price>55.8</price>
            </book>
            <book>
                <bookName>呼啸山庄</bookName>
                <price>99</price>
            </book>
        }} {{ student} 王二狗 30 male {{ books}  [{{ book} 十万个为啥 22.8} {{ book} 从入门到放弃 68}]
            <book>
                <bookName>十万个为啥</bookName>
                <price>22.8</price>
            </book>
            <book>
                <bookName>从入门到放弃</bookName>
                <price>68</price>
            </book>
        }}]
    <student>
        <studentName>韩茹</studentName>
        <age>31</age>
        <sex>female</sex>
        <books>
            <book>
                <bookName>红与黑</bookName>
                <price>55.8</price>
            </book>
            <book>
                <bookName>呼啸山庄</bookName>
                <price>99</price>
            </book>
        </books>
    </student>
    <student>
        <studentName>王二狗</studentName>
        <age>30</age>
        <sex>male</sex>
        <books>
            <book>
                <bookName>十万个为啥</bookName>
                <price>22.8</price>
            </book>
            <book>
                <bookName>从入门到放弃</bookName>
                <price>68</price>
            </book>
        </books>
    </student>
}
*/
