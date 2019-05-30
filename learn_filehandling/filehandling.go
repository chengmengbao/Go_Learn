package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	// "os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 拉取更新测试
func main() {
	// box := packr.NewBox("../filehandling")
	// data := box.String("test.txt")
	// fmt.Println("Contents of file:", data)
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	// // fmt.Println(*fptr)
	// data, err := ioutil.ReadFile(*fptr)
	// check(err)
	// fmt.Printf("%s", string(data))
	// data, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println("文件读取错误：", err)
	// 	return
	// }
	// fmt.Println("1 文件的内容：", data)
	// fmt.Printf("data type:%T\n", data)
	// fmt.Println("2 文件的内容：", string(data))

	// write_data := []byte("hello go I interest you\n")
	// err1 := ioutil.WriteFile("test.txt", write_data, 0644)
	// if err1 != nil {
	// 	fmt.Println("文件写入错误：", err)
	// 	return
	// }
	// f, err := os.Open("test.txt")
	// check(err)

	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// o2, err := f.Seek(6, 0)
	// check(err)
	// b2 := make([]byte, 2)
	// n2, err := f.Read(b2)
	// check(err)
	// fmt.Printf("%d bytes @ %d :%s\n", n2, o2, string(b2))

}
