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

	ok, err := conn.Do("SET", "name", "chengmengbao")
	if err != nil {
		fmt.Println("redis set error: ", err)
	} else {
		fmt.Printf("SET name: %s\n", ok) //SET name: OK
	}

	//ok, err = conn.Do("expire", "name", 10) //10秒过期
	//if err != nil {
	//	fmt.Println("set expire error: ", err)
	//	return
	//}

	//time.Sleep(time.Second * 9) //延迟10秒

	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Get name:%s \n", name) //Get name:chengmengbao
	}

}
