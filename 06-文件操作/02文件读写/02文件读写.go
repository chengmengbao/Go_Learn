package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件，即新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("新建文件出错")
		return
	}
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
		// fmt.Println("n = ", n)
	}

	defer f.Close()
}

func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*2)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("读取文件失败")
		return
	}
	fmt.Println("buf=", string(buf[:n]))
}

func ReadFileline(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = ###%s###\n", string(buf))
	}
}
func main() {
	path := "./demo.txt"
	// WriteFile(path)
	// ReadFile(path)
	ReadFileline(path)
}
