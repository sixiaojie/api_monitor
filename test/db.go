package test
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/op/go-logging"
)


func JDBC(log  *logging.Logger) (db *sql.DB, result bool){
	db_info := ConfigParser("db")
	host := db_info["host"]
	port := db_info["port"]
	database := db_info["database"]
	user := db_info["user"]
	password := db_info["password"]
	charset := db_info["charset"]
	connect_S := user+":"+password+"@tcp"+"("+host+":"+port+")"+"/"+database+"?"+"charset="+charset
	db,err := sql.Open("mysql",connect_S)
	if err != nil {
		log.Error(err)
		db = nil
		return nil,false
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	return db,true
}

