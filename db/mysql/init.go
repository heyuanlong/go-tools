package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

func NewMysql(user, password, ip, port, mysqldb string) (*sql.DB, error) {

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		user,
		password,
		ip,
		port,
		mysqldb,
	)
	MysqlClient, err := sql.Open("mysql", addr)
	if err != nil {
		log.Println("mysql open fail:", err)
		return nil, err
	}
	MysqlClient.SetMaxOpenConns(2000)
	MysqlClient.SetMaxIdleConns(10)
	MysqlClient.Ping()
	return MysqlClient, nil
}

// var (
// 	MysqlClient     *sql.DB
// )
// func InitMysql()  {

// 	user,_ := 		kconf.GetString("mysql","user")
// 	password,_ := 	kconf.GetString("mysql","password")
// 	ip,_ := 		kconf.GetString("mysql","ip")
// 	port,_ := 		kconf.GetString("mysql","port")
// 	mysqldb,_ := 	kconf.GetString("mysql","mysqldb")

// 	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
// 		user,
// 		password,
// 		ip,
// 		port,
// 		mysqldb,
// 	)
// 	MysqlClient, _ = sql.Open("mysql", addr)
// 	MysqlClient.SetMaxOpenConns(2000)
// 	MysqlClient.SetMaxIdleConns(10)
// 	MysqlClient.Ping()
// 	klog.Warn.Printf("mysql open ok")
// }
