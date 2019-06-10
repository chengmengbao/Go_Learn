package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

//操作示例，示例中将使用两个goroutine分别担任发布者和订阅者角色进行演示：
//发布、订阅

//订阅者
func Subs() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	psc := redis.PubSubConn{conn}
	psc.Subscribe("channel1") //订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

//发布者
func Push(message string) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}

	//defer conn.Close()  //注意不要加这句，不然发布者就退出了

	_, err1 := conn.Do("PUBLISH", "channel1", message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}
}

func main() {
	go Subs()
	go Push("I am WinerChan")
	time.Sleep(time.Second * 3)
}

//输出结果
//channel1: subscribe 1
//channel1: message: I am WinerChan
