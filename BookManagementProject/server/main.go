package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"main.go/Mysql"
)

func main() {

	start()
	//Mysql.Textbook()
	//Mysql.AddBook()
}

func start() {
	//中文
	r := gin.Default()
	var sql Mysql.Mysql
	sql.StartMysql()
	defer sql.EndMysql()
	user := r.Group("/user")

	//登录
	user.GET("/login", sql.Login)
	//注册
	user.GET("/regist", sql.Regist)
	//修改密码
	user.GET("/modify", sql.Modify)
	//获取信息
	user.GET("/getinformation", sql.GetInformation)
	//添加记录
	user.GET("/addrecord", sql.AddRecord)
	//书籍
	book := r.Group("/book")
	//获取书籍
	book.GET("/getbook", sql.GetBook)
	//获取章节
	book.GET("/getchapter", sql.GetChapter)
	//获取排行榜
	book.GET("/getrank", sql.GetRank)
	//搜索书籍
	book.GET("/searchbook", sql.SearchBook)
	//添加评论
	book.GET("/addcomment", sql.AddComment)
	//获取分类
	book.GET("/getbooksbytype", sql.GetBooksByCategory)

	r.Run(":8080")
}
