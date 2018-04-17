package test

import (
	"strings"
)

/*
这里简单的检查Linux中目录，文件。根据字符串返回路径，文件，已经是否有错误
 */

func GetDirFile(path string) (dirname,file string, err bool) {
	dir_path := ""
	if path == ""{
		return "","",false
	}
	s := strings.Split(path,"/")
	if len(s) == 2 {
		return "",s[1],true
	}else {
		for i:=1;i <len(s)-1;i++ {
			dir_path = dir_path +"/"+s[i]
		}
		return dir_path,s[len(s)-1],true
	}

}




