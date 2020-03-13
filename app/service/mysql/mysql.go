package mysql

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
)

func GetVersion() (gdb.Record, error) {
	db := g.DB()
	versionInfo, err := db.GetOne("select version()")
	return versionInfo, err
}

func IsExistTable(tableName string) bool {
	db := g.DB()
	res, err := db.GetOne("SHOW TABLES LIKE '" + tableName + "'")
	if err != nil {
		panic(err)
	}

	if len(res.Map()) > 0 {
		return true
	}

	return false
}

func CreateTable(tableName string) bool {
	var sql []byte
	var err error
	switch {
	case tableName == "siam_logs":
		sql, err = ioutil.ReadFile("./public/resource/siam_logs.sql")
		if err != nil {
			fmt.Printf("打开文件失败" + err.Error())
			return false
		}

	case tableName == "siam_projects":
		sql, err = ioutil.ReadFile("./public/resource/siam_projects.sql")
		if err != nil {
			fmt.Printf("打开文件失败" + err.Error())
			return false
		}

	case tableName == "siam_api_log":
		sql, err = ioutil.ReadFile("./public/resource/siam_api_log.sql")
		if err != nil {
			fmt.Printf("打开文件失败" + err.Error())
			return false
		}

	}
	db := g.DB()

	_, execErr := db.Exec(string(sql))
	if execErr != nil {
		return false
	}
	return true
}
