package Mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DATABSAE           = "book_management"
	TABLE_USER         = "user_information"
	ID                 = "id"
	NAME               = "name"
	PASSWORD           = "_password"
	QQ                 = "qq"
	PHONE              = "phone"
	EMAIL              = "email"
	INTRODUCE          = "introduction"
	LOOK_RECORD        = "look_record"
	REGEST_TIME        = "regest_time"
	LASTEST_LOGIN_TIME = "lastest_login_time"
	AGE                = "age"
	RECORDS            = "record"
)
const (
	LOGIN          = "login"
	REGIST         = "regist"
	MODIFY         = "modify"
	GET_INFOMATION = "get information"
)

type Mysql struct {
	db *sql.DB
}

const (
	NOT_EXIST      = "not exist"
	SUCCESS        = "success"
	PASSWORD_ERROR = "password error"
	USER_EXIST     = "user exist"
	REGIST_SUCCESS = "regist success"
	REGIST_ERROR   = "regist error"

	MODIFY_ERROR   = "modify error"
	MODIFY_SUCCESS = "modify success"
)

type UserInformationMysql struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	Password           string   `json:"password"`
	Qq                 string   `json:"qq"`
	Phone              string   `json:"phone"`
	Email              string   `json:"email"`
	Introduce          string   `json:"introduce"`
	Look_record        string   `json:"look_record"`
	Regest_time        []uint8  `json:"regist_time"`
	Lastest_login_time []uint8  `json:"lastest_login_time"`
	Age                string   `json:"age"`
	Records            []record `json:"records"`
}

type comment struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	Create_time []uint8   `json:"Create_time"`
	User        string    `json:"user"`
	Comments    []comment `json:"child_Comments"`
	Like_count  int       `json:"like_id"`
}
type Book struct {
	Id             int       `json:"id"`
	Create_time    []uint8   `json:"create_time"`
	Update_time    []uint8   `json:"update_time"`
	Author         string    `json:"author"`
	Chapters_count int       `json:"chapters_count"`
	Like_count     int       `json:"like_count"`
	Record_count   int       `json:"record_count"`
	Comments       []comment `json:"comment"`
	Name           string    `json:"name"`
	Hit            int       `json:"Hit"`
	Chapters       []chapter `json:"chapters"`
	Look_count     int       `json:"look_count"`
	Introduce      string    `json:"introduce"`
	Category       string    `json:"category"`
}

type chapter struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type chapter_content struct {
	Id           int     `json:"id"`
	Book_id      int     `json:"book_id"`
	Chapter_name string  `json:"chapter_name"`
	Chapter_old  int     `json:"chapter_old"`
	Create_time  []uint8 `json:"create_time"`
	Update_time  []uint8 `json:"update_time"`

	Content string `json:"content"`
}

func (e *Mysql) StartMysql() error {
	var err error
	e.db, err = sql.Open("mysql", "coffeerain:vaxwk4f6pEETwsET@tcp(mysql5.sqlpub.com:3310)/huaweibookmanagement")
	if err != nil {
		log.Fatalln("mysql start error", err)
	} else {
		log.Println("mysql start success")
	}
	log.Println("成功连接数据库 huaweibookmanagement")
	return err
}
func (e *Mysql) GetInformationName(name string, user *UserInformationMysql) error {
	st, err := e.db.Prepare("SELECT * FROM user_information WHERE " + NAME + "=?")
	if err != nil {
		log.Println("查询失败NAME", err)
		return err
	}
	defer st.Close()
	row, err := st.Query(name)
	if err != nil {
		log.Println("查无此人:", err)
		return err
	}
	defer row.Close()
	for row.Next() {
		var str string
		err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Qq, &user.Phone, &user.Email, &user.Introduce, &user.Look_record, &user.Regest_time, &user.Lastest_login_time, &user.Age, &str)
		if err != nil {
			log.Println("赋值失败:", err)
			return err
		}
		err = json.Unmarshal([]byte(str), &user.Records)
		if err != nil {
			log.Println("赋值失败:", err)
			return err
		}
	}
	return err
}
func (e *Mysql) GetInformation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user UserInformationMysql

	err := e._GetInformationId(id, &user)
	if err != nil {

		SeedMessage(c, GET_INFOMATION, NOT_EXIST, err)
		return
	}
	SeedMessage(c, GET_INFOMATION, SUCCESS, user)
	return
}
func (e *Mysql) _GetInformationId(id int, user *UserInformationMysql) error {

	st, err := e.db.Prepare("SELECT * FROM user_information WHERE " + ID + "=?")
	if err != nil {
		log.Println("查询失败")
	}
	defer st.Close()
	row, err := st.Query(id)
	if err != nil || row.Err() != nil {
		log.Println("查询失败")
		return err
	}
	defer row.Close()
	for row.Next() {
		var str string
		err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Qq, &user.Phone, &user.Email, &user.Introduce, &user.Look_record, &user.Regest_time, &user.Lastest_login_time, &user.Age, &str)
		if err != nil {
			log.Println("赋值失败:", err)
			return err
		}
		err = json.Unmarshal([]byte(str), &user.Records)
		if err != nil {
			log.Println("赋值失败:", err)
			return err
		}
	}

	return err
}
func (e *Mysql) Modify(c *gin.Context) {
	column := c.Query("column")
	value := c.Query("value")
	id := c.Query("id")
	_, err := e.db.Query("UPDATE " + TABLE_USER + " SET " + column + " = \"" + value + "\" WHERE " + ID + " = " + id)
	if err != nil {
		SeedMessage(c, MODIFY, MODIFY_ERROR, err)

		return
	}
	SeedMessage(c, MODIFY, MODIFY_SUCCESS, err)

	return
}

func (e *Mysql) Regist(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	qq := c.Query("qq")
	phone := c.Query("phone")
	email := c.Query("email")
	introduce := c.Query("introduce")
	look_record := c.Query("look_record")
	age := c.Query("age")
	if age == "" {
		age = "20"
	}
	if look_record == "" {
		look_record = "无"
	}
	// 确保必填字段不为空
	if name == "" || password == "" {
		SeedMessage(c, REGIST, REGIST_ERROR, "用户名和密码不能为空")
		return
	}

	var records []record
	records = append(records, record{
		Bookname: "初始书籍",
		Time:     time.Now().Round(time.Second),
	})

	bit, err := json.Marshal(records)
	if err != nil {
		SeedMessage(c, REGIST, REGIST_ERROR, err)
		return
	}

	if introduce == "" {
		introduce = "暂无介绍"
	}
	if qq == "" {
		qq = "0"
	}
	if phone == "" {
		phone = "0"
	}
	if email == "" {
		email = "无"
	}

	_, err = e.db.Exec(`
		INSERT INTO huaweibookmanagement.user_information 
		(name, _password, qq, phone, email, introduction, look_record, age,record) 
		VALUES (?, ?, ?, ?, ?, ?, ?,?,?)`,
		name, password, qq, phone, email, introduce, look_record, age, bit)

	if err != nil {
		SeedMessage(c, REGIST, REGIST_ERROR, err)
		log.Println("注册错误:", err)
		return
	}

	SeedMessage(c, REGIST, REGIST_SUCCESS, "注册成功")
}
func (e *Mysql) EndMysql() error {
	err := e.db.Close()
	if err != nil {
		log.Fatalln("mysql close error", err)
	}
	return err

}

type StatusMessage struct {
	Time   time.Time   `json:"time"`
	Type   string      `json:"type"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SeedMessage(c *gin.Context, _type string, status string, data interface{}) {
	bit, err := json.Marshal(&StatusMessage{
		Time:   time.Now(),
		Type:   _type,
		Status: status,
		Data:   data,
	})
	if err != nil {
		log.Println("json.Marshal err:", err)
	}
	c.Data(200, "application/json", bit)
}

func (e *Mysql) Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	var userInform UserInformationMysql
	err := e.GetInformationName(name, &userInform)

	if err != nil {
		if err == sql.ErrNoRows {
			// 明确用户不存在的情况
			SeedMessage(c, LOGIN, NOT_EXIST, "用户不存在")
		} else {
			// 其他数据库错误
			log.Printf("登录查询错误: %v", err)
			SeedMessage(c, LOGIN, NOT_EXIST, "系统错误，请稍后再试")
		}
		return
	}
	if userInform.Password == password {
		//response, _ := json.Marshal(userInform)
		SeedMessage(c, LOGIN, SUCCESS, userInform)

		return
	} else {
		SeedMessage(c, LOGIN, PASSWORD_ERROR, nil)

		return
	}

}

type record struct {
	Bookname string    `json:"booKname"`
	Time     time.Time `json:"time"`
}

func (e *Mysql) AddRecord(c *gin.Context) {
	bookname := c.Query("bookname")
	id, _ := strconv.Atoi(c.Query("id"))
	var user UserInformationMysql
	e._GetInformationId(id, &user)
	fmt.Println(user)
	user.Records = append(user.Records, record{
		Bookname: bookname,
		Time:     time.Now()})
	err := e._Modify(user)
	if err == nil {
		SeedMessage(c, "add record", SUCCESS, user)
	}

}
func (e *Mysql) _Modify(new UserInformationMysql) error {
	//bit, err := json.Marshal(new.Records)

	bit, err := json.Marshal(new.Records)
	if err != nil {
		log.Println("json.Marshal err:", err)
		return err
	}
	_, err = e.db.Exec("UPDATE user_information SET name = ?, _password = ?, qq = ?, phone = ?, email = ?,introduction = ?,look_record = ?,lastest_login_time = ?,age = ?,record = ? WHERE id = ?",
		new.Name, new.Password, new.Qq, new.Phone, new.Email, new.Introduce, new.Look_record, new.Lastest_login_time, new.Age, bit, new.Id)
	if err != nil {
		log.Println("修改错误:" + err.Error())
		return err
	}
	return err
}
func (e *Mysql) _GetBookInformation(name string, book *Book) error {
	query := `
		SELECT id, create_time, update_time, author, chapters_count, like_count, record_count,
		       comments, book_name, hit, chapters, introduce
		FROM books WHERE book_name = ? LIMIT 1`

	row := e.db.QueryRow(query, name)
	var co, ca string
	err := row.Scan(&book.Id, &book.Create_time, &book.Update_time, &book.Author,
		&book.Chapters_count, &book.Like_count, &book.Record_count,
		&co, &book.Name, &book.Hit, &ca, &book.Introduce)
	if err != nil {
		log.Println("扫描数据失败:", err)
		return err
	}
	_ = json.Unmarshal([]byte(co), &book.Comments)
	_ = json.Unmarshal([]byte(ca), &book.Chapters)
	return nil
}

func (e *Mysql) _ModifyBook(new *Book) error {
	bit, err := json.Marshal(new.Comments)
	ca, err := json.Marshal(new.Chapters)
	_, err = e.db.Exec("UPDATE books SET update_time = ?, chapters_count = ?, like_count = ?, record_count = ?, comments = ?, book_name = ?, look_count = ?, chapters = ? WHERE id = ?;",
		new.Update_time, new.Chapters_count, new.Like_count, new.Record_count, bit, new.Name, new.Look_count, ca, new.Id)

	if err != nil {
		log.Println("修改错误")
		return err
	}
	return err
}
func (e *Mysql) _AddBook(book *Book) error {
	bit, err := json.Marshal(book.Comments)

	ca, err := json.Marshal(book.Chapters)
	_, err = e.db.Exec("INSERT INTO books (create_name,chapters_count,like_count,record_count,comments,book_name,look_count,chapters) VALUES (?,?,?,?,?,?,?,?)",
		book.Author, book.Chapters_count, book.Like_count, book.Record_count, bit, book.Name, book.Look_count, ca)
	if err != nil {
		log.Println("book add error")

	}
	return err

}

// type chapter_content struct {
// 	Id           int     `json:"id"`
// 	Book_id      int     `json:"book_id"`
// 	Chapter_name string  `json:"chapter_name"`
// 	Create_time  []uint8 `json:"create_time"`
// 	Update_time  []uint8 `json:"update_time"`
// 	Author       string  `json:"author"`
// 	Content      string  `json:"content"`
// }

func (e *Mysql) _AddChapter(book *Book, ch *chapter_content) error {
	ch.Book_id = book.Id
	ch.Id = book.Id*10000 + ch.Chapter_old
	_, err := e.db.Exec("INSERT INTO chapters (id,book_id,chapter_name,chapters_order,content) VALUES (?,?,?,?,?)",
		ch.Id, ch.Book_id, ch.Chapter_name, ch.Chapter_old, ch.Content)
	if err != nil {
		log.Println("book add error")
		log.Println(err)
	} else {

		book.Chapters_count = ch.Chapter_old

		book.Chapters = append(book.Chapters, chapter{
			Id:   ch.Id,
			Name: ch.Chapter_name,
		})
		e._ModifyBook(book)
	}
	return err
}
func (e *Mysql) GetChapter(c *gin.Context) {
	_id := c.Query("id")
	id, _ := strconv.Atoi(_id)
	ch := new(chapter_content)
	err := e._GetChapter(id, ch)
	if err != nil || ch.Chapter_name == "" {
		SeedMessage(c, "get chapter", "error", nil)
		return

	} else {
		SeedMessage(c, "get chapter", "success", ch)
	}
}

func (e *Mysql) _GetChapter(id int, ch *chapter_content) error {
	st, err := e.db.Prepare("SELECT * FROM chapters WHERE " + ID + "=?")
	if err != nil {
		log.Println("查询失败")
		log.Println(err)
		return err
	}
	defer st.Close()
	row, err := st.Query(id)
	if err != nil || row.Err() != nil {
		log.Println("查询失败")
		log.Println(err)
		return err
	}
	defer row.Close()

	for row.Next() {
		var str string
		err = row.Scan(&ch.Id, &ch.Book_id, &ch.Chapter_name, &ch.Chapter_old, &ch.Create_time, &ch.Update_time, &str, &ch.Content)
		if err != nil {
			log.Println("查询失败")
			log.Println(err)
			return err
		}
	}

	return err
}
func (e *Mysql) _ModifyChapters(ch *chapter_content) error {
	_, err := e.db.Exec("UPDATE chapters SET book_id = ?, caphter_name = ?, chapters_older = ?, update_time = ?, content = ? WHERE id = ?;",
		ch.Chapter_name, ch.Chapter_old, ch.Update_time, ch.Content, ch.Id)
	if err != nil {
		log.Println("修改错误")
		return err
	}
	return err
}
func (e *Mysql) Delete(id int, table string) error {
	_, err := e.db.Exec("DELETE FROM "+table+" WHERE id = ?;", id)
	if err != nil {
		log.Println("delete error")
	}
	return err
}

func (e *Mysql) _GetBookInformationID(id int, book *Book) error {
	st, err := e.db.Prepare(`
		SELECT id, create_time, update_time, author, chapters_count, like_count, record_count,
		       comments, book_name, hit, chapters, introduce
		FROM books WHERE id = ?`)
	if err != nil {
		log.Println("查询失败")
		return err
	}
	defer st.Close()

	row, err := st.Query(id)
	if err != nil || row.Err() != nil {
		log.Println("查询失败")
		return err
	}
	defer row.Close()

	for row.Next() {
		var co, ca string
		err = row.Scan(&book.Id, &book.Create_time, &book.Update_time, &book.Author,
			&book.Chapters_count, &book.Like_count, &book.Record_count,
			&co, &book.Name, &book.Hit, &ca, &book.Introduce)
		if err != nil {
			log.Println("查询失败:", err)
			return err
		}
		_ = json.Unmarshal([]byte(co), &book.Comments)
		_ = json.Unmarshal([]byte(ca), &book.Chapters)
	}
	return nil
}

func (e *Mysql) GetBook(c *gin.Context) {
	count := c.Query("count")
	if count == "" {
		count = "15"
	}
	// count, _ := strconv.Atoi(_count)
	var books []Book

	st, err := e.db.Query(`
  SELECT id, create_time, update_time, author, chapters_count, like_count, record_count,
         comments, book_name, hit, chapters, introduce
  FROM books ORDER BY RAND() LIMIT ` + count + `;`)

	if err != nil {
		log.Println("获取失败")
		log.Println(err)
		SeedMessage(c, "get books", "error", nil)

	}

	defer st.Close()

	for st.Next() {
		var book Book
		var co, ca string
		err = st.Scan(&book.Id, &book.Create_time, &book.Update_time, &book.Author,
			&book.Chapters_count, &book.Like_count, &book.Record_count,
			&co, &book.Name, &book.Hit, &ca, &book.Introduce)

		println(co)
		println(ca)
		err = json.Unmarshal([]byte(co), &book.Comments)
		err = json.Unmarshal([]byte(ca), &book.Chapters)

		books = append(books, book)
	}

	SeedMessage(c, "get books", "success", books)

}
func Textbook() {
	var sql Mysql
	sql.StartMysql()
	defer sql.EndMysql()
	var book Book
	book.Id = 1
	book.Author = "聂伟山"
	book.Name = "Go语言"
	book.Look_count = 100
	c := comment{
		Id:          1,
		Content:     "123",
		Create_time: []uint8{1, 2, 3},
		User:        "sdf",
		Comments:    []comment{},
		Like_count:  1,
	}

	book.Comments = append(book.Comments, c)
	err := sql._GetBookInformation("go语言编程", &book)
	fmt.Println(book, err)
	log.Println(err)

}
func AddBook() {
	change := []string{"斗破", "重启", "传奇", "斗罗大陆", "斗罗大陆2", "斗罗大陆3", "斗罗大陆4", "斗罗大陆5", "斗罗大陆6", "斗罗大陆7", "斗罗大陆8", "斗罗大陆9", "斗罗大陆10", "斗罗大陆11", "斗罗大陆12", "斗罗大陆13", "斗罗大陆14", "斗罗大陆15", "斗罗大陆16", "斗罗大陆17", "斗罗大陆18", "斗罗大陆19", "斗罗大陆20", "斗罗大陆21", "斗罗大陆22", "斗罗大陆23", "斗罗大陆24", "斗罗大陆25", "斗罗大陆26", "斗罗大陆27", "斗罗大陆28", "斗罗大陆29", "斗罗大陆30", "斗罗大陆31", "斗罗大陆32", "斗罗大陆33", "斗罗大陆34", "斗罗大陆35", "斗罗大陆36", "斗罗大陆37", "斗罗大陆38", "斗罗大陆39", "斗罗大陆40", "斗罗大陆41", "斗罗大陆42", "斗罗大陆43", "斗罗大陆44", "斗罗大陆45", "斗罗大陆46", "斗罗大陆47", "斗罗大陆48", "斗罗大陆49", "斗罗大陆50", "斗罗大陆51", "斗罗大陆52", "斗罗大陆53", "斗罗大陆54", "斗罗大陆55", "斗罗大陆56", "斗罗大陆57", "斗罗大陆58", "斗罗大陆59", "斗罗大陆60", "斗罗大陆61", "斗罗大陆62", "斗罗大陆63", "斗罗大陆64", "斗罗大陆65", "斗罗大陆66", "斗罗大陆67", "斗罗大陆68", "斗罗大陆69", "斗罗大陆70", "斗罗大陆71", "斗罗大陆72", "斗罗大陆73", "斗罗大陆74", "斗罗大陆75", "斗罗大陆76", "斗罗大陆77", "斗罗大陆78", "斗罗大陆79", "斗罗大陆80", "斗罗大陆81", "斗罗大陆82", "斗罗大陆83", "斗罗大陆84", "斗罗大陆85", "斗罗大陆86", "斗罗大陆87", "斗罗大陆88", "斗罗大陆89", "斗罗大陆90", "斗罗大陆91", "斗罗大陆92", "斗罗大陆93", "斗罗大陆94", "斗罗大陆95", "斗罗大陆96", "斗罗大陆97", "斗罗大陆98", "斗罗大陆99", "斗罗大陆100", "斗罗大陆101", "斗罗大陆102", "斗罗大陆103", "斗罗大陆104", "斗罗大陆105", "斗罗大陆106", "斗罗大陆107", "斗罗大陆108", "斗罗大陆109", "斗罗大陆110", "斗罗大陆111", "斗罗大陆112", "斗罗大陆113", "斗罗大陆114", "斗破苍穹1", "斗破苍穹2"}
	//first:="每一个钟头上传一章，直到传完二十章！红票和收藏别忘了～）北凉王府龙盘虎踞于清凉山，千门万户，极土木之盛。作为王朝硕果仅存的异姓王，在庙堂和江湖都是毁誉参半的北凉王徐骁作为一名功勋武臣，可谓得到了皇帝宝座以外所有的东西，在西北三州，他就是当之无愧的主宰，只手遮天，翻云覆雨。难怪朝廷中与这位异姓王政见不合的大人们私下都会文绉绉骂一声徐蛮子，而一些居心叵测的，更诛心地丢了顶“二皇帝”的帽子。今天王府很热闹，位高权重的北凉王亲自开了中门，摆开辉煌仪仗，迎接一位仙风道骨的老者，府中下人们只听说是来自道教圣地龙虎山的神仙，相中了痴痴傻傻的小王爷，要收作闭关弟子，这可是天大的福缘，北凉王府都解释成傻人有傻福。可不是，小王爷自打出生起便没哭过，读书识字一窍不通，六岁才会说话，名字倒是威武气派，徐龙象，传闻还是龙虎山的老神仙当年给取的，说好十二年后再来收徒，这不就如约而至了。王府内一处院落，龙虎山师祖一级的道门老祖宗捻着一缕雪白胡须，眉头紧皱，背负一柄不常见的小钟馗式桃木剑，配合他的相貌，确实当得出尘二字，谁看都要由衷赞一声世外高人呐。但此番收徒显然遇到了不小的阻碍，倒不是王府方面有异议，而是他的未来徒弟犟脾气上来了，蹲在一株梨树下，用屁股对付他这个天下道统中论地位能排前三甲的便宜师傅，至于武功嘛，咳咳，前三十总该有的吧。连堂堂大柱国北凉王都得蹲在那里好言相劝，循循善诱里透着股诱拐，“儿子，去龙虎山学成一身本事，以后谁再敢说你傻，你就揍他，三品以下的文官武将，打死都不怕，爹给你撑腰。”“儿啊，你力气大，不学武捞个天下十大高手当当就太可惜了。学成归来，爹就给你一个上骑都尉当当，骑五花马，披重甲，多气派。”小王爷完全不搭理，死死盯着地面，瞧得津津有味。“黄蛮儿，你不是喜欢吃糖葫芦吗，那龙虎山遍地的野山楂，你随便摘随便啃。赵天师，是不是？”老神仙硬挤出一抹笑容，连连点头称是。收徒弟收到这份上，也忒寒碜了，说出去还不被全天下笑话。可哪怕位于堂堂超一品官职、在十二郡一言九鼎的大柱国口干舌燥了，少年还是没什么反应，估计是不耐烦了嫌老爹说得呱噪，翘起屁股，噗一下来了个响屁，还不忘扭头对老爹咧嘴一笑。把北凉王给气得抬手作势要打，可抬着手僵持一会儿，就作罢。一来是不舍得打，二来是打了没意义。这儿子可真对得起名字，徐龙象，取自“水行中龙力最大，陆行中象力第一，威猛如金刚，是谓龙象”，别看绰号黄蛮儿的傻儿子憨憨笨笨，至今斗大字不识，皮肤病态的暗黄，身形比较同龄人都要瘦弱，但这气力，却是一等一骇人。徐骁十岁从军杀人，从东北锦州杀匈奴到南部灭大小六国屠七十余城再到西南镇压蛮"
	//B2:="没有见过，但如小儿子这般可天生铜筋铁骨力拔山河的，真没有。徐骁心中轻轻叹息，黄蛮儿若能稍稍聪慧一些，心窍多开一二，将来必定可以成为陷阵第一的无双猛将啊。他缓缓起身转头朝龙虎山辈分极高的道士尴尬一笑，后者眼神示意不打紧，只是心中难免悲凉，收个徒弟收到这份上，也忒不是个事儿了，一旦传出去还不得被天下人笑话，这张老脸就甭想在龙虎山那一大帮徒子徒孙面前摆放喽。束手无策的北凉王心生一计，嘿嘿道：“黄蛮儿，你哥游行归来，看时辰也约莫进城了，你不出去看看？”小王爷猛地抬头，表情千年不变的呆板僵硬，但寻常木讷无神的眼眸却爆绽出罕见光彩，很刺人，拉住老爹的手就往外冲。可惜这北凉王府出了名百廊回转曲径千折，否则也容不下一座饱受朝廷清官士大夫们诟病的“听潮亭”，手被儿子握得生疼的徐骁不得不数次提醒走错路了，足足走了一炷香时间，这才来到府外。父子和老神仙身后，跟着一帮扛着大小箱子的奴仆，都是准备带往龙虎山的东西，北凉王富可敌国，对儿女也是素来宠溺，见不得他们吃一点苦受一点委屈。到了府外，小王爷一看到街道空荡，哪里有哥哥的身影，先是失望，继而愤怒，沉沉嘶吼一声，沙哑而暴躁，起先想对徐骁发火，但笨归笨，起码还知道这位是父亲，否则徐骁的下场恐怕就得像前不久秋狩里倒霉遇到徐龙象的黑罴了，被单枪匹马的十二岁少年生生撕成两半。他怒瞪了一眼心虚的老爹，掉头就走。不希望功亏一篑的徐骁无奈丢给老神仙一个眼神。龙虎山真人微微一笑，伸出枯竹一般的手臂，但仅是两指搭住了小王爷的手腕，轻声慈祥道：“徐龙象，莫要浪费了你百年难遇的天赋异禀，随我去龙虎山，最多十年，你便可下山立功立德。”少年也不废话，哼了一声，继续前往，但玄妙古怪的是他发现自己没能挣脱老道士看似云淡风轻的束缚，那踏出去悬空的一步如何都没能落地。北凉王如释重负，这位道统辈分高到离谱的上人果真还是有些本事的，知子莫若父，徐骁哪里不知道小儿子的力道，霸气得很，以至于他都不敢多安排仆人女婢给儿子，生怕一个不小心就捏断了胳膊腿脚，这些年院中被坐坏拍烂的桌椅不计其数，也亏得北凉王府家底厚实，寻常殷实人家早就破产了。小王爷愣了一下，随即发火，轻喝一声，硬是带着老神仙往前走了一步，两步，三步。头顶黄冠、身披道袍的真人只是微微咦了一声，不怒反喜，悄悄加重了几分力道，阻止了少年的继续前行。如此一来，徐龙象是真怒了，面容狰狞如同一只野兽，伸出空闲的一只手，双手握住老道士的手臂，双脚一沉，咔嚓，在白玉地板上踩出两个坑，一甩，就将老道士整个人给丢掷了出去。大柱国徐骁眯起眼睛，丝毫不怕惹出命案，那道士若没这个斤两本事，摔死就摔死好了，他徐骁连不可一世的西楚王朝都给用凉州铁骑踏平了，何时对江湖门派有过丝毫的敬畏？天下道统首领龙虎山又如何？所辖境内数个大门大派虽比不上龙虎山，但在王朝内也属一流规模，例如那数百年一直跟龙虎山争那道统的武当山，在江湖上够超然了吧，还不是每年都主动派人送来三四炉珍品丹药？老道士轻轻飘荡到王府门口的一座两人高汉白玉石狮子上，极富仙人气势。光凭这一手，若是搁在市井中，那还不得搏得满堂喝彩啊。这按照北凉王世子即徐骁嫡长子的那个脍炙人口的说法，那就是“该赏，这活儿不简单，是技术活”，指不定就是几百几千银票打赏出去了，想当年世子殿下还没出北凉祸害别人的时日，多少青楼清伶或者江湖骗子得了他的阔绰赏钱。最高纪录是一位外地游侠，在街上一言不合与当地剑客相斗，从街边菜摊打起打到湖畔最后打到湖边凉州最大鹞子溢香楼的楼顶，把白日宣--淫的世子给吵醒了，立马顾不得白嫩如羊脂美玉的花魁小娘子，在窗口大声叫好，事后在世子殿下的掺和下官府非但没有追究，反而差点给那名游侠送去凉州好男儿的大锦牌，他更是让仆人快马加鞭送去一大摞整整十万银票。没有喜好玩鹰斗犬的世子殿下的大好陵州，可真是寂寞啊。正经人家的小娘们终于敢漂漂亮亮上街买胭脂了，二流纨绔们终于没了跟他们抢着欺男霸女的魔头了，大大小小的青楼也等不到那位头号公子哥的一掷千金了。北凉王徐骁生有二女二子，俱是奇葩。大郡主出嫁，连克三位丈夫，成了王朝内脸蛋最俏嫁妆最多的寡妇，在江南道五郡艳名远播，作风放浪。二郡主虽相貌平平，却是博学多才，精于经纬，师从上阴学宫韩谷子韩大家，成了兵法大家许煌、纵横术士司马灿等一干帝国名流的小师妹。徐龙象是北凉王的最小儿子，相对声名不显，而大儿子则是连京城那边都有大名声的家伙，一提起大柱国徐骁，必然会扯上世子徐凤年，“赞誉”一声虎父无犬子，可惜徐骁是英勇在战场上，儿子却是争气在风花雪月的败家上。三年前，世子殿下徐凤年传言被脖子上架着刀剑撵出了王府，被迫去学行关中豪族年轻后辈及冠礼之前的例行游历，一晃就是三载，彻底没了音信，陵州至今记得世子殿下出城时，城墙上十几号大纨绔和几十号大小花魁眼中含泪的感人画面，只是有内幕说等世子殿下走远了，当天，红雀楼的酒宴便通了个宵，太多美酒倒入河内，整座城都闻得见酒香。回到王府这边，心窍闭塞的小王爷奔跑冲向玉石狮子，似乎摔一个老头子不过瘾，这次是要把碍眼的老道连同号称千钧重的狮子一同摔出去。只是他刚摇晃起狮子，龙虎山老道便飘下了来，牵住少年的一只手，使出真功夫，以道门晦涩的“搬山”手法，巧妙一带，就将屈膝半蹲的少年拉起身，轻笑道：“黄蛮儿，不要闹，随为师去吧。”少年一只手握住狮子底座边角，五指如钩，深入玉石，不肯松手，双臂拉伸如猿猴，嘶哑嚷着：“我要等哥哥回来，哥哥说要给我带回天下第一美女做媳妇，我要等他！”位极人臣的大柱国徐骁哭笑不得，无可奈何，望向黄冠老道，重重叹气道：“罢了，再等等吧，反正也快了。”老道士闻言，笑容古怪，但还是松开了小王爷的手臂，心中咂舌，这小家伙何止是天生神力，根本就是太白星下凡嘛。不过，那个叫徐凤年的小王八蛋真的要回来了？这可不是一个好消息。想当年他头回来王府，可是吃足了"
	//b3:="不说，那才七八岁的兔崽子直接放了一群恶犬来咬自己，后来好不容易解释清楚，进了府邸，小王八玩意就又坏心眼了，派了两位娇滴滴的美娇--娘三更半夜来敲门，说是天气冷要暖被子，若非贫道定力超凡脱俗，还真就着了道，现在偶尔想起来，挺后悔没跟两位姑娘彻夜畅聊《大洞真经》和《黄庭经》，即便不聊这个，聊聊《素女心经》也好嘛。黄昏中，官道上一老一少被余晖拉长了身影，老的背负着一个被破布包裹的长条状行囊，衣衫褴褛，一头白发，还夹杂几根茅草，弄个破碗蹲地上就能乞讨了，牵着一匹瘦骨嶙嶙的跛马。小的其实岁数不小，满脸胡茬，一身市井麻衫，逃荒的难民一般。“老黄，再撑会儿，进了城回了家，就有大块肉大碗酒了，他娘的，以前没觉得这酒肉是啥稀罕东西，现在一想到就嘴馋得不行，每天做梦都想。”瞧不出真实年龄的年轻男人有气没力道。仆人模样的邋遢老头子呵呵一笑，露出一口缺了门牙的黄牙，显得贼憨厚贼可笑。“笑你个大爷，老子现在连哭都哭不出来了。”年轻人翻白眼道，他是真没那个精神气折腾了。两千里归途，就只差没落魄到沿路乞讨，这一路下水里摸过鱼，上山跟兔子捉迷藏，爬树掏过鸟窝，只要带点荤的，弄熟了，别管有没有盐巴，那就都是天底下最美味的一顿饭了。期间经过村庄试图偷点鸡鸭啥的，好几次被扛锄头木棍的壮汉追着跑了几十里路，差点没累死。哪个膏粱子弟不是鲜衣怒马威风八面？再瞧瞧自个儿，一袭破烂麻衣，草鞋一双，跛马一只，还不舍得宰了吃肉，连骑都不舍得，倒是多了张蹭饭的嘴。恶奴就更没有了，老黄这活了一甲子的小身板他光是瞅着就心慌，生怕这行走两千里路哪天就没声没息嗝屁了，到时候他连个说话的伴儿都没有，还得花力气在荒郊野岭挖个坑。尚未进城，城墙外头不远有一个挂杏花酒的摊子，他实在是精疲力尽了，闻着酒香，闭上眼睛，抽了抽鼻子，一脸陶醉，真贼娘的香。一发狠，他走过去寻了一条唯一空着的凳子一屁股坐下，咬牙使出最后气力喊道：“小二，上酒！”身边出城或者进城中途歇息的酒客都嫌弃这衣着寒碜的一主一仆，刻意坐远了。生意忙碌的店小二原本听着声音要附和一声“好嘞”，可一看主仆两人的装束，立即就拉下脸，出来做买卖的，没个眼力劲儿怎么样，这两位客人可不想是掏得出酒钱的货色，店小二还算厚道，没立马赶人，只是端着皮笑肉不笑的笑脸提醒道：“我们这招牌杏花酒可要一壶二十钱，不贵，可也不便宜。”若是以前，被如此狗眼看人低，年轻人早就放狗放恶奴了，可三年世态炎凉，过习惯了身无分文的日子，架子脾气收敛了太多，喘着气道：“没事，自然有人来结账，少不了你的打赏钱。”“打赏？”店小二扯开了嗓门，一脸鄙夷。年轻人苦笑，拇指食指放在嘴边，把最后那点吃奶的力气都使出来吹了一声哨子，然后就趴在简陋酒桌上，打鼾，竟然睡着了。店小二只觉得莫名其妙，唯有眼尖的人依稀瞧见头顶闪过一点影子。一头鹰隼般的飞禽如箭矢掠过城头。大概酒客喝光一碗杏花酒的时光，大地毫无征兆地轰鸣起来，酒桌摇晃，酒客们瞪大眼睛看着酒水跟着木桌一起晃荡，都小心翼翼捧起来，四处张望。只见城门处冲出一群铁骑，绵延成两条黑线，仿佛没个尽头。尘土飞扬中，高头大马，俱是北凉境内以一当百名动天下的重甲骁骑，看那为首扛旗将军手中所拿的王旗，鲜艳如血，上书一字，“徐”！乖乖，北凉王麾下的嫡系军。天下间，谁能与驰骋辗转过王朝南北十三州的北凉铁骑争锋？以往，西楚王朝觉得它的十二万大戟士敢逆其锋芒，可结果呢，景河一战，全军覆没，降卒悉数坑杀，哀嚎如雷。两百精锐铁骑冲刺而出，浩浩荡荡，气势如虹。头顶一只充满灵气的鹰隼似在领路。两百铁骑瞬间静止，动作如出一辙，这份娴熟，已经远远超出一般行伍悍卒百战之兵的范畴。正四品武将折冲都尉翻身下马，一眼看见牵马老仆，立即奔驰到酒肆前，跪下行礼，恭声道：“末将齐当国参见世子殿下！”而那位口出狂言要给打赏钱的寒酸年轻人只是在睡梦中呢喃了一句，“小二，上酒。”"
	file, _ := os.ReadFile("D:\\王俊林-py作业\\雪中悍刀行.txt")

	str := string(file)
	var con []string
	for i := 0; i < len(str)-1024; i += 1024 {
		con = append(con, string(str[i:i+1024]))
	}

	var sql Mysql
	sql.StartMysql()
	defer sql.EndMysql()
	for i := 65; i < len(change); i++ {
		rand.Seed(time.Now().UnixNano())

		// 生成一个0到100之间的随机整数
		num := rand.Intn(1000)

		var book Book
		book.Id = i + 15
		book.Author = "聂伟山"
		book.Name = change[i]
		book.Look_count = num * i

		err := sql._AddBook(&book)
		if err != nil {
			log.Println(err)
		}
		// //
		err = sql._GetBookInformation(book.Name, &book)
		if err != nil {
			log.Println(err)
		}

		for j := 1; j < 100; j++ {
			for k := 0; ; k++ {
				var chapter chapter_content
				chapter.Chapter_name = "第" + strconv.Itoa(j) + "章"
				chapter.Chapter_old = j
				chapter.Content = con[(j+i+k)%255]
				err := sql._AddChapter(&book, &chapter)
				if err == nil {
					break
				}
			}

		}

	}
}
func (e *Mysql) SearchBook(c *gin.Context) {
	bookName := c.Query("bookname")
	var books []Book

	query := `
		SELECT id, create_time, update_time, author, chapters_count, like_count, record_count,
		       comments, book_name, hit, chapters, introduce
		FROM books WHERE book_name LIKE ? ORDER BY hit DESC`
	param := "%" + bookName + "%"

	rows, err := e.db.Query(query, param)
	if err != nil {
		log.Println("搜索书籍失败:", err)
		SeedMessage(c, "search book", "error", nil)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		var co, ca string
		err = rows.Scan(&book.Id, &book.Create_time, &book.Update_time, &book.Author,
			&book.Chapters_count, &book.Like_count, &book.Record_count,
			&co, &book.Name, &book.Hit, &ca, &book.Introduce)
		_ = json.Unmarshal([]byte(co), &book.Comments)
		_ = json.Unmarshal([]byte(ca), &book.Chapters)
		books = append(books, book)
	}
	SeedMessage(c, "search book", "success", books)
}

func (e *Mysql) GetRank(c *gin.Context) {
	count := c.Query("count")
	if count == "" {
		count = "10"
	}
	var books []Book

	st, err := e.db.Query(`
		SELECT id, create_time, update_time, author, chapters_count, like_count, record_count,
		       comments, book_name, hit, chapters, introduce
		FROM books ORDER BY hit DESC LIMIT ` + count + `;`)
	if err != nil {
		log.Println("获取失败")
		SeedMessage(c, "get rank", "error", nil)
		return
	}
	defer st.Close()

	for st.Next() {
		var book Book
		var co, ca string
		err = st.Scan(&book.Id, &book.Create_time, &book.Update_time, &book.Author,
			&book.Chapters_count, &book.Like_count, &book.Record_count,
			&co, &book.Name, &book.Hit, &ca, &book.Introduce)
		_ = json.Unmarshal([]byte(co), &book.Comments)
		_ = json.Unmarshal([]byte(ca), &book.Chapters)
		books = append(books, book)
	}
	SeedMessage(c, "get rank", "success", books)
}

func (e *Mysql) AddComment(c *gin.Context) {
	bit := c.Query("bit")
	var book Book
	err := json.Unmarshal([]byte(bit), &book)
	if err != nil {
		log.Println("获取失败")
		return
	}
	e._ModifyBook(&book)

}

func (e *Mysql) Text() {
	var userInform UserInformationMysql
	e.GetInformationName("123", &userInform)

}

func (e *Mysql) GetBooksByCategory(c *gin.Context) {
	category := c.Query("type")
	if category == "" {
		SeedMessage(c, "get books by category", "error", "分类不能为空")
		return
	}

	query := "SELECT * FROM books WHERE category = ?"
	rows, err := e.db.Query(query, category)
	if err != nil {
		log.Println("按分类获取书籍失败:", err)
		SeedMessage(c, "get books by category", "error", nil)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		var co, ca string
		err = rows.Scan(&book.Id, &book.Create_time, &book.Update_time, &book.Author,
			&book.Chapters_count, &book.Like_count, &book.Record_count, &co,
			&book.Name, &book.Hit, &ca, &book.Category)
		if err != nil {
			log.Println("扫描书籍数据失败:", err)
			continue
		}
		_ = json.Unmarshal([]byte(co), &book.Comments)
		_ = json.Unmarshal([]byte(ca), &book.Chapters)
		books = append(books, book)
	}
	SeedMessage(c, "get books by category", "success", books)
}
