package gb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 数据库连接
var DB *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "51!Bim_!@#"
	NETWORK  = "tcp"
	SERVER   = "120.78.198.120"
	PORT     = 3306
	DATABASE = "monitor"
)

// InitDb 初始化数据库
func InitDb() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		Logger.Error("Open mysql failed,err:%v\n", err)
		return
	}
	err = db.Ping() //尝试连接校验用户名密码
	if err != nil {
		Logger.Error("open %s failed, err:%s\n", dsn, err)
	}
	DB = db
	fmt.Printf("Db connected success !\n")
}

// SaveToStockDay 保存数据到 stock_day数据库
func SaveToDB(msg, ip string) {
	sqlStr := fmt.Sprintf(`insert into device_data(msg,ip)values('%s','%s')`, msg, ip)
	if DB == nil {
		InitDb()
		fmt.Println("DB is nil")
		return
	}
	_, err := DB.Exec(sqlStr)
	if err != nil {
		Logger.Error("insert faild,err:%v\n", err)
		fmt.Printf("insert faild,err:%v\n", err)
	}
}
