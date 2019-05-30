package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//获取命令行参数
	listAgrs := os.Args
	if len(listAgrs) != 3 {
		fmt.Println("命令行参数必须需要三个")
		return
	}

	srcFilename := listAgrs[1]
	dstFilename := listAgrs[2]
	if srcFilename == dstFilename {
		fmt.Println("源文件不能和目的文件相同")
		return
	}

	//只读方式打开源文件
	sF, err := os.Open(srcFilename)
	if err != nil {
		fmt.Println("打开源文件失败，err = ", err)
		return
	}

	//新建一个目的文件
	dF, err := os.Create(dstFilename)
	if err != nil {
		fmt.Println("新建源文件失败，err = ", err)
		return
	}

	//函数结束后，关闭文件
	defer sF.Close()
	defer dF.Close()

	buf := make([]byte, 4*1024) //创建临时缓冲区，4k大小
	for {
		n, err := sF.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
		}

		dF.Write(buf[:n])

	}

}
