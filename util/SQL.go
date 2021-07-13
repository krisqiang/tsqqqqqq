package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type MysqlConfig struct {
	Ip           string
	Port         int64
	Username     string
	Password     string
	DatabaseName string
	Url          string
}

func init() {
	mysqlConfig := new(MysqlConfig)
	mysqlConfig.Ip = "localhost"
	mysqlConfig.Port = 3306
	mysqlConfig.Username = "root"
	mysqlConfig.Password = "root"
	mysqlConfig.DatabaseName = "tsqqqqqq"
	db = mysqlConfig.Connect()
}

func (mysql *MysqlConfig) Connect() *sql.DB {
	mysql.build()
	db, err := sql.Open("mysql", mysql.Url)
	if err != nil {
		panic(err)
	}
	return db
}

func (mysql *MysqlConfig) build() {
	//TODO format is username:password@tcp(ip:port)/databaseName?charset=utf8
	mysql.Url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysql.Username, mysql.Password, mysql.Ip, mysql.Port, mysql.DatabaseName)
	fmt.Println(mysql.Url)
}

func Query(sql string, obj ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(sql, obj...)
	return rows, err
}

func ExecData(sql string, obj []interface{}) (res sql.Result, err error) {
	fmt.Println("Exec", sql, cap(obj))
	res, err = db.Exec(sql, obj...)
	return
}
