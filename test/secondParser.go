package test

import (
	"net/url"
	"errors"
	"os"
	"fmt"
	"time"
	"strconv"
)

func CreateTable(){
	sql := "create table if not exists task(id int primary key auto_increment,url varchar(20) not null,cycle int default 60,result varchar(20) default 'ok',time int)"
	err := Sql_Exec(sql)
	if err != nil {
		logger.Error(err)
		os.Exit(10)
	}
}

func Sql_Exec(sql string) (err error){
	if sql ==""{
		logger.Warning("sql语句不不能为空")
		return errors.New("sql语句不不能为空")
	}else {
		res,err := db.Exec(sql)
		if err != nil {
			logger.Error(err)
			return errors.New("新建监控失败")
		}
		id,err := res.LastInsertId()
		if err != nil {
			logger.Warning(sql,"没插入到数据库,id未变：%d",id,err.Error())
			return errors.New("没插入到数据库")
		}
		return nil
	}
}

func ParserValue(value url.Values) (err error){
	url := value.Get("url")
	if url == "" {
		return errors.New("需要填写url")
	}
	err = check_url_repeat(url)
	if err != nil {
		return err
	}
	cycle := value.Get("cycle")
	if cycle == ""{
		cycle = "60"
	}
	//time,err := strconv.ParseInt(times, 10, 64)
	result := value.Get("result")
	if result == ""{
		result = "ok"
	}
	unix_time := strconv.FormatInt(time.Now().Unix(),10)
	sql := "insert into task(url,cycle,result,time) values('"+url+"',"+cycle+",'"+result+"',"+unix_time+")"
	fmt.Println(sql)
	err = Sql_Exec(sql)
	return err
}
