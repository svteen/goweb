package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	DbName     = "sf2"
	DbUser     = "root"
	DbPass     = "1111"
	DbType     = "mysql"
	OriginTime = "2006-01-02 15:04:05"
)

type User struct {
	Id         int64  `orm:"pk;auto"`
	Username   string `orm:"size(60)"`
	Password   string `orm:"size(32)"`
	Email      string `orm:"size(100)"`
	Status     int
	Nicknanme  string `orm:"size(60)"`
	Avatar     string `orm:size(100)`
	Registered time.Time
}

type Category struct {
	Id        int `orm:"pk;auto"`
	Parent_id int
	Name      string `orm:"size(60)"`
	Created   time.Time
	Updated   time.Time
}

type Post struct {
	Id      int `orm:"pk;auto"`
	User_id int64
	Cate_id int
	Title   string `orm:"size(100)"`
	Excerpt string `orm:"size(250)"`
	Content string `orm:"text"`
	Author  string `orm:"size(60)"`
	Posted  time.Time
	Created time.Time
	Updated time.Time
	Status  int
}

type Prize struct {
	Id         int `orm:"pk;auto"`
	Lottery_id int
	Name       string `orm:size(100)`
	Qty        int
	Status     int
	Sort_order int
}

type Lottery struct {
	Id       int64 `orm:"pk;auto"`
	Prize_id int
	Name     string `orm:size(100)`
	Started  int
	Ended    int
	Chance   int
	Status   int
}

type Record struct {
	Id         int `orm:"pk;auto"`
	User_id    int
	Lottery_id int
	Prize_id   int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Category), new(Post), new(Prize), new(Lottery), new(Record))
}

func GetUserInfo(id int64) *User {
	o := orm.NewOrm()
	user := User{Id: id}
	o.Read(&user)
	return &user
}

func UpdateUserInfo(id int64, nickname, email, avatar string) bool {
	o := orm.NewOrm()
	user := User{Id: id, Nicknanme: nickname, Email: email, Avatar: avatar}
	id, err := o.Update(&user, "Nicknanme", "Email", "Avatar")
	if err == nil {
		return true
	}
	return false
}

func GetPosts() *Post {
	o := orm.NewOrm()
	post := Post{}
	o.Read(&post)
	return &post
}

func GetUserByName(username string) *User {
	o := orm.NewOrm()
	user := User{Username: username}
	o.Read(&user, "username")
	return &user
}

func SavePost(title, excerpt, origin, content string) int64 {
	o := orm.NewOrm()
	now := time.Now()
	post := Post{Cate_id: 0,
		Title:   title,
		Excerpt: excerpt,
		Author:  origin,
		Content: content,
		//2006-01-02 15:04:05
		Created: now,
		Posted:  time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		Updated: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		Status:  0}
	id, err := o.Insert(&post)
	if err == nil {
		return id
	}
	return 0
}

func UpdatePost(pid int, title, excerpt, origin, content string) int64 {
	o := orm.NewOrm()
	now := time.Now()
	post := Post{Id: pid,
		Cate_id: 0,
		Title:   title,
		Excerpt: excerpt,
		Author:  origin,
		Content: content,
		Updated: now}
	id, err := o.Update(&post, "Id", "Title", "Excerpt", "Author", "Content", "Updated")
	beego.Debug("error-------------->", err)
	if err == nil {
		return id
	}
	return 0
}

func SaveCategory(parent_id int, name string) int64 {
	o := orm.NewOrm()
	now := time.Now()
	category := Category{Parent_id: parent_id,
		Name:    name,
		Created: now,
		Updated: now}
	id, err := o.Insert(&category)
	if err == nil {
		return id
	}
	return 0
}

func GetPostNums() int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("post").Count()
	if err == nil {
		return cnt
	}
	return 0
}

func GetCategoryNums() int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("category").Count()
	if err == nil {
		return cnt
	}
	return 0
}

func GetPaginationPost(limit int, offset int) *[]Post {
	o := orm.NewOrm()
	qs := o.QueryTable("post")
	post := new([]Post)
	qs.OrderBy("-Created").Limit(limit, offset).All(post)
	beego.Debug("postFunc")
	return post
}

func GetPaginationCate(limit int, offset int) *[]Category {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	category := new([]Category)
	qs.OrderBy("-Created").Limit(limit, offset).All(category)
	beego.Debug("postFunc")
	return category
}

func GetArticle(id int) *Post {
	o := orm.NewOrm()
	post := Post{Id: id}
	o.Read(&post)
	return &post
}

func GetCategory(id int) *Category {
	o := orm.NewOrm()
	category := Category{Id: id}
	o.Read(&category)
	return &category
}

func DelCategory(id int) int64 {
	o := orm.NewOrm()
	category := Category{Id: id}
	num, _ := o.Delete(&category)
	return num
}

func DelArticle(id int) int64 {
	o := orm.NewOrm()
	post := Post{Id: id}
	num, _ := o.Delete(&post)
	return num
}

func GetAppointedArticle(userid int64, id int) *Post {
	o := orm.NewOrm()
	post := Post{User_id: userid, Id: id}
	o.Read(&post)
	return &post
}

func GetAllCategory() *[]Category {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	category := new([]Category)
	qs.OrderBy("-Id").All(category)
	return category
}
