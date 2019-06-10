package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}

	defer conn.Close()

	conn.Send("HMSET", "user", "name", "hanru", "age", "30")
	conn.Send("HSET", "user", "sex", "female")
	conn.Send("HGET", "user", "sex")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
	//输出结果：
	//Receive res1:OK
	//Receive res2:0
	//Receive res3:female
}
