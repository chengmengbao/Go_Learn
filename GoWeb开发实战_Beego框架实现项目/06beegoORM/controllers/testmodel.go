package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ModelController struct {
	beego.Controller
}

//-----------定义struct-------------
type User struct {
	// 对应user表
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

/**
由于model定义为UserInfo，那么实际上操作的表示：user_info
*/
type UserInfo struct {
	Id       int64  `orm:"column(id)"`       // 也可以省略不写，orm会自动映射
	Username string `orm:"column(username)"` // 也可以省略不写，orm会自动映射
	Password string
}

type MyUser struct {
	Id   int
	Name string
	Age  int
}

func init() {
	//注册驱动：如果是默认的三个可以不写
	orm.RegisterDriver("mysql", orm.DRMySQL) //可以不加

	//注册默认数据库，ORM 必须注册一个别名为default的数据库，作为默认使用。
	/*
		参数一：数据库别名
		参数二：驱动名称
		参数三：数据库连接字符串:username:password@tcp(127.0.0.1:3306)/databasename?charset=utf8
		参数四：设置数据库的最大空闲连接
	*/
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/go_tmpdb?charset=utf8", 30)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag), new(UserInfo))

}
func (c *ModelController) CreateTable() {
	//自动建表
	orm.RunSyncdb("default", false, true)
	datainit()
}

func datainit() {
	o := orm.NewOrm()
	//rel  : 自动生成外键为 表名_id
	sql1 := "insert into user (name,profile_id) values ('hanru',1),('ruby',2),('王二狗',3);"
	sql2 := "insert into profile (age) values (20),(19),(21);"
	sql3 := "insert into tag (name) values ('offical'),('beta'),('dev');"
	sql4 := "insert into post (title,user_id) values ('paper1',1),('paper2',1),('paper3',2),('paper4',3),('paper5',3);"
	// m2m 生成的 表名：子表_主表s  主键自增
	sql5 := "insert into post_tags (tag_id, post_id) values (1,1),(1,3),(2,2),(3,3),(2,4),(3,4),(3,5); "

	//使用Raw（）.Exec（）执行sql
	o.Raw(sql1).Exec()
	o.Raw(sql2).Exec()
	o.Raw(sql3).Exec()
	o.Raw(sql4).Exec()
	o.Raw(sql5).Exec()
}

func (c *ModelController) Get() {
	o := orm.NewOrm()
	o.Using("default") // 可以省略不写。你可以使用Using函数指定其他数据库

	/**
	  通过orm对象来进行数据库的操作，这种情况是必须要知道主键
	*/

	// 1. insert
	//user := UserInfo{Username:"张三1",Password:"zhangsan1231"}
	//id, err := o.Insert(&user)

	//2. update
	//user := UserInfo{Id:2, Username:"lisi1", Password:"lisi123"}
	//num, err := o.Update(&user) //第一个返回值为影响的行数

	//3. delete
	//num, err := o.Delete(&UserInfo{Id: 2})
	//if err != nil {
	//	fmt.Println("err = ", err)
	//}

	// 4.read ,查询
	//user := UserInfo{Id:1}
	//err := o.Read(&user)
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	c.Ctx.WriteString(fmt.Sprintf("id:%d, username:%s, password:%s\n", user.Id, user.Username, user.Password))
	//}

	//5.ReadOrCreate, 查询或创建
	user := UserInfo{Username: "李小花", Password: "xiaohuazzzz"}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if created, id, err := o.ReadOrCreate(&user, "username", "password"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)

		} else {
			fmt.Println("Get an object. Id:", id)
			fmt.Printf("id:%d, username:%s, password:%s\n", user.Id, user.Username, user.Password)
		}
		c.Ctx.WriteString(fmt.Sprintf("created:%t, id:%d", created, id))
	} else {
		fmt.Println("err = ", err)
	}
}

//高级查询
func (c *ModelController) Query() {
	orm.Debug = true //是否开启调试模式，调试模式下回打印出sql
	o := orm.NewOrm()
	o.Using("default") // 可以省略不写。你可以使用Using函数指定其他数据库

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("user")

	// 也可以直接使用对象作为表名
	//user := new(User)
	//qs := o.QueryTable(user) // 返回 QuerySeter
	// 2.指定查询：
	//qs.Filter("name", "hanru") // WHERE name = 'hanru'
	//qs.Filter("name__exact", "hanru") // WHERE name = 'hanru'
	//qs.Filter("name__iexact", "hanru")
	// WHERE name LIKE 'hanru'
	// 大小写不敏感，匹配任意 'Han' 'hAN'
	// 使用 = 匹配，大小写是否敏感取决于数据表使用的 collation
	//qs.Filter("profile_id", nil) // WHERE profile_id IS NULL

	//qs.Filter("name__contains", "hanru")
	// WHERE name LIKE BINARY '%hanru%'
	// 大小写敏感, 匹配包含 hanru 的字符

	//qs.Filter("name__icontains", "hanru")
	// WHERE name LIKE '%hanru%'
	// 大小写不敏感, 匹配任意 'im Hanru', 'im hanRu'

	//qs.Filter("profile__age__in", 17, 18, 19, 20)
	// WHERE profile.age IN (17, 18, 19, 20)
	//ids:=[]int{17,18,19,20}
	//qs.Filter("profile__age__in", ids)
	// WHERE profile.age IN (17, 18, 19, 20)

	//qs.Filter("profile__age__gt", 17)
	// WHERE profile.age > 17
	//qs.Filter("profile__age__gte", 18)
	// WHERE profile.age >= 18

	//qs.Filter("profile__age__lt", 17)
	// WHERE profile.age < 17
	//qs.Filter("profile__age__lte", 18)
	// WHERE profile.age <= 18

	//qs.Filter("name__startswith", "hanru")
	// WHERE name LIKE BINARY 'hanru%'
	// 大小写敏感, 匹配以 'hanru' 起始的字符串
	//qs.Filter("name__istartswith", "hanru")
	// WHERE name LIKE 'hanru%'
	// 大小写不敏感, 匹配任意以 'hanru', 'Hanru' 起始的字符串

	//qs.Filter("name__endswith", "hanru")
	// WHERE name LIKE BINARY '%hanru'
	// 大小写敏感, 匹配以 'hanru' 结束的字符串
	//qs.Filter("name__iendswithi", "hanru")
	// WHERE name LIKE '%hanru'
	// 大小写不敏感, 匹配任意以 'hanru', 'Hanru' 结束的字符串

	//qs.Filter("profile__isnull", true)
	//qs.Filter("profile_id__isnull", true)
	// WHERE profile_id IS NULL
	//qs.Filter("profile__isnull", false)
	// WHERE profile_id IS NOT NULL

	// 多个 Filter 之间使用 AND 连接
	//qs.Filter("profile__isnull", true).Filter("name", "hanru")
	// WHERE profile_id IS NULL AND name = 'hanru'

	// 使用 NOT 排除条件
	// 多个 Exclude 之间使用 AND 连接
	//qs.Exclude("profile__isnull", true).Filter("name", "hanru")
	// WHERE NOT profile_id IS NULL AND name = 'hanru'

	// 在 expr 前使用减号 - 表示 DESC 的排列
	//qs.OrderBy("id", "-profile__age")
	// ORDER BY id ASC, profile.age DESC
	//qs.OrderBy("-profile__age", "profile")
	// ORDER BY profile.age DESC, profile_id ASC

	// 对应 sql 的 distinct 语句, 返回不重复的值.
	//qs.Distinct()
	// SELECT DISTINCT

	// 依据当前的查询条件，返回结果行数
	//cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
	//fmt.Printf("Count Num: %s, %s", cnt, err)

	//exist := o.QueryTable("user").Filter("UserName", "Name").Exist()
	//fmt.Printf("Is Exist: %s", exist)

	// 依据当前查询条件，进行批量更新操作
	//num, err := o.QueryTable("user").Filter("name", "hanru").Update(orm.Params{
	//	"name": "ruby",
	//})
	//fmt.Printf("Affected Num: %s, %s", num, err)
	// SET name = "ruby" WHERE name = "hanru"

	// 依据当前查询条件，进行批量删除操作
	//num, err := o.QueryTable("user").Filter("name", "hanru").Delete()
	//fmt.Printf("Affected Num: %s, %s", num, err)
	// DELETE FROM user WHERE name = "hanru"

	//var users []*User
	//num, err := o.QueryTable("user").Filter("name", "hanru").All(&users)
	//fmt.Printf("Returned Rows Num: %s, %s", num, err)

	// 尝试返回单条记录
	//var user User
	//err := o.QueryTable("user").Filter("name", "hanru").One(&user)
	//if err == orm.ErrMultiRows {
	//	// 多条的时候报错
	//	fmt.Printf("Returned Multi Rows Not One")
	//}
	//if err == orm.ErrNoRows {
	//	// 没有找到记录
	//	fmt.Printf("Not row found")
	//}

	//定义一个User类型的切片
	var users []*User
	//num, err := qs.All(&users)
	num, err := qs.Filter("profile__age__in", 17, 18, 19, 20).All(&users)
	if err != nil {
		// 处理err
		fmt.Println("qs.All() err = ", err)
	}
	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")
	for _, user := range users {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Name) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Profile.Id) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")
}

//原生查询
func (c *ModelController) QuerySQL() {
	//是否开启调试模式，调试模式下回打印出sql
	orm.Debug = true
	o := orm.NewOrm()
	// 可以省略不写。你可以使用Using函数指定其他数据库
	o.Using("default")

	/* 1.查询
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM user_info").Values(&maps)
	if err != nil {
		//处理err
		fmt.Println("o.Raw err = ", err)
	}

	fmt.Printf("Result Nums: %d\n", num)
	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, m := range maps {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", m["id"]) + "</td>" +
			"<td>" + fmt.Sprintf("%v", m["username"]) + "</td>" +
			"<td>" + fmt.Sprintf("%v", m["password"]) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")
	*/
	// 2. exec()
	/*
		res, err := o.Raw("UPDATE user_info SET username = ? where id = ?", "尼古拉斯", "3").Exec()
		if err != nil {
			//处理err
			c.Ctx.WriteString(fmt.Sprintf("err:%v", err))
		}
		num, _ := res.RowsAffected()
		c.Ctx.WriteString(fmt.Sprintf("共影响了 num:%d 条数据。。", num))
	*/

	// 3. QueryRow
	var user UserInfo
	err := o.Raw("SELECT id, username, password FROM user_info WHERE id = ?", 1).QueryRow(&user)
	if err != nil {
		// 处理err
		c.Ctx.WriteString(fmt.Sprintf("err:%v", err))
	}
	c.Ctx.WriteString(fmt.Sprintf("id:%d, username:%s, password:%s", user.Id, user.Username, user.Password))
}

// QueryBuilder查询
func (c *ModelController) QueryBuilder() {

	orm.Debug = true //是否开启调试模式，调试模式下回打印出sql
	var users []MyUser

	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("user.id",
		"user.name",
		"profile.age").
		From("user").
		InnerJoin("profile").On("user.profile_id = profile.id").
		Where("age > ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	num, err := o.Raw(sql, 19).QueryRows(&users)
	if err != nil {
		// 处理err
	}

	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, user := range users {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Name) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Age) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")

}
