package test

import (
	"errors"
)

func check_url_repeat(url string) (err error){
	rows,err := db.Query("select count(*) from task where url = ?",url)
	if err != nil {
		return errors.New("查看是否有重复url遇见错误")
	}
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			logger.Error(err)
		}
	}
	if count >= 1 {
		return errors.New("该url已经重复，请注意")
	}
	return nil
}
