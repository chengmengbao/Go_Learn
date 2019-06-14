### 一、使用Redis数据库
=====
#### 1.1 安装
golang操作redis的客户端包有多个比如redigo、go-redis，github上Star最多的莫属redigo。

github地址：https://github.com/garyburd/redigo

目前已经迁移到：https://github.com/gomodule/redigo

文档：https://godoc.org/github.com/garyburd/redigo/redis

我们可以直接通过go get来安装它。

```
go get github.com/garyburd/redigo/redis
```


#### 1.2 Pipelining(管道)
管道操作可以理解为并发操作，并通过Send()，Flush()，Receive()三个方法实现。客户端可以使用send()方法一次性向服务器发送一个或多个命令，
命令发送完毕时，使用flush()方法将缓冲区的命令输入一次性发送到服务器，
客户端再使用Receive()方法依次按照先进先出的顺序读取所有命令操作结果。

```$xslt
Send(commandName string, args ...interface{}) error
Flush() error
Receive() (reply interface{}, err error)
```

- Send：发送命令至缓冲区
- Flush：清空缓冲区，将命令一次性发送至服务器
- Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。


#### 1.3 发布/订阅
redis本身具有发布订阅的功能，其发布订阅功能通过命令SUBSCRIBE(订阅)／PUBLISH(发布)实现，
并且发布订阅模式可以是多对多模式还可支持正则表达式，发布者可以向一个或多个频道发送消息，
订阅者可订阅一个或者多个频道接受消息。

示意图：

发布者：

![Image text](https://raw.githubusercontent.com/chengmengbao/img-folder/master/mark.png?token=AEIEPIAXQK2L7VOHOKE3GLC47J6XO)

订阅者：

![Image text](https://raw.githubusercontent.com/chengmengbao/img-folder/master/mark1.png?token=AEIEPIC5D353KJ3MEE5H7ZK47J7GQ)

#### 1.4 事务操作
MULTI, EXEC,DISCARD和WATCH是构成Redis事务的基础，当然我们使用go语言对redis进行事务操作的时候本质也是使用这些命令。

MULTI：开启事务

EXEC：执行事务

DISCARD：取消事务

WATCH：监视事务中的键变化，一旦有改变则取消事务。
