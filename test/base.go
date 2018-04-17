package test
import (
	"github.com/larspensjo/config"
	"os"
	"fmt"
)

var configFile = "/Users/sijie/go/src/logdb/conf/default.ini"

/*
这里输出的格式是类似Python中字典的格式，格式里面包含着配置文件，及配置信息。如log来说
格式就是{"log":"","level":""}
 */
func ConfigParser(item string) (map[string]string){
	pro := make(map[string]string)
	c, err := config.ReadDefault(configFile)
	checkError(err)
	if c.HasSection(item) == false {
		fmt.Printf("不存在%s配置",item)
		os.Exit(10)
	}
	section,err := c.SectionOptions(item)
	if err == nil {
		for _,v := range section {
			options,err := c.String(item,v)
			if err == nil {
				pro[v] = options
			}
		}
	}
	return pro
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
	return
}

