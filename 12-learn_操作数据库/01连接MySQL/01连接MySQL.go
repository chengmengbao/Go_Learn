package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "password"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go_tmpdb"
)

//Db数据库连接池
var DB *sql.DB

type User struct {
	c1 int
	c2 string
	c3 string
}

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

//insert
func InsertUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO tmptab (`c1`, `c2`, `c3`) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.c1, user.c2, user.c3)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	////获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

//delete
func DeleteUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM tmptab WHERE c1=?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.c1)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

//update
func UpdateUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE tmptab SET c2 = ?, c3 = ? WHERE c1 = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.c2, user.c3, user.c1)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

//query
func SelectUserById(c1 int) User {
	var user User
	err := DB.QueryRow("SELECT * FROM tmptab WHERE c1 = ?", c1).Scan(&user.c1, &user.c2, &user.c3)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return user
}

func SelectAllUser() []User {
	//执行查询语句
	rows, err := DB.Query("SELECT * from tmptab")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var users []User
	//循环读取结果
	for rows.Next() {
		var user User
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.c1, &user.c2, &user.c3)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}

func main() {
	//初始化
	InitDB()
	//增
	//user := User{3, "chengmengbao3", "password3"}
	//InsertUser(user)

	//删
	//user = User{c1: 1}
	//DeleteUser(user)

	//改
	//UpdateUser(user)

	//查
	//user = SelectUserById(1)
	//fmt.Println("user=", user)

	//多行查询
	users := SelectAllUser()
	fmt.Println("users=", users)
}
