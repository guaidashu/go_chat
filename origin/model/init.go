package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go_chat/app/libs"
)

var DbEngine *xorm.Engine

func init() {

	var err error

	err = getConnection()

	if err != nil {
		fmt.Println(err.Error())
	}

}

func getConnection() (err error) {

	engine, err := xorm.NewEngine("mysql", "yy:wyysdsa!@(127.0.0.1:3306)/go_chat/charset=utf8")

	if err != nil {
		return libs.NewReportError(err)
	}

	DbEngine = engine

	DbEngine.ShowSQL(true)

	// 设置 数据库最大打开的连接数
	DbEngine.SetMaxOpenConns(10)

	libs.DebugPrint("xorm初始化 数据库成功")

	return

}

func createTable(m BaseModel) {

	if ok, err := m.IsExists(); !ok {

		if err != nil {
			libs.DebugPrint(fmt.Sprintf("%v", libs.NewReportError(err)))
		}

		if err = m.CreateTable(); err != nil {
			libs.DebugPrint(fmt.Sprintf("%v", libs.NewReportError(err)))
		}

	}

}

func CreateTable() {

	userModel := new(UserModel)

	createTable(userModel)

}
