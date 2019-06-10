package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

//hash操作
func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("HMSET", "user", "name", "Paul", "age", 14)
	if err != nil {
		fmt.Println("redis HSET error:", err)
	}

	res, err := redis.Int64(conn.Do("HGET", "user", "age"))
	if err != nil {
		fmt.Println("redis HGET error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %d \n", res)
	}
	//输出结果：
	//res type : int64
	//res  : 14
}
