package test

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)



var format = logging.MustStringFormatter(`%{time:15:04:05.000} %{shortfunc} > %{level:.4s}  %{message}`,)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}




func LoggerSetting(section string) *logging.Logger{
	var log = logging.MustGetLogger("example")
	var file map[string]string = ConfigParser(section)
	logfile := file["logfile"]
	path,_,result := GetDirFile(logfile)
	//fmt.Println(logfile,path)
	if result != true {
		fmt.Println("请配置配置文件中日志路径")
		os.Exit(10)
	}
	_,err := os.Stat(path)
	if err != nil {
		//fmt.Println("发现错误，请确定后再次启动，错误原因为:",err)
		//os.Exit(10)
		os.MkdirAll(path,0644)
	}

	logFile, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	if err != nil{
		fmt.Println(err)
		os.Exit(10)
	}
	backend := logging.NewLogBackend(logFile, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backend2Formatter)
	return log
}



