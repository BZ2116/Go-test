//数据库好久没用了，不知道为什么打不开了。。。
//我随便贴一个数据库初始化在这，，，，，

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

import (
	"database/sql"
)


var db *sql.DB


func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:1347715183@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True"
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	} else {
		fmt.Printf("连接成功")
	}
}
