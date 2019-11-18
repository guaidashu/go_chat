package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go_chat/app/libs"
	"xorm.io/core"
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

	engine, err := xorm.NewEngine("mysql", "yy:wyysdsa!@(127.0.0.1:3306)/go_chat?charset=utf8")

	if err != nil {
		return libs.NewReportError(err)
	}

	DbEngine = engine

	// shows sql statement on standard output or your io.Writer
	DbEngine.ShowSQL(true)

	// show debug and other informations
	DbEngine.Logger().SetLevel(core.LOG_DEBUG)

	// 设置 数据库最大打开的连接数
	DbEngine.SetMaxOpenConns(10)

	libs.DebugPrint("xorm初始化 数据库成功")

	return

}

func createTable(m ...BaseModel) {

	for _, model := range m {
		if ok, err := model.IsExists(); !ok {

			if err != nil {
				libs.DebugPrint(fmt.Sprintf("%v", libs.NewReportError(err)))
			}

			if err = model.CreateTable(); err != nil {
				libs.DebugPrint(fmt.Sprintf("%v", libs.NewReportError(err)))
			}

		}
	}

}

func CreateTable() {

	userModel := new(UserModel)

	contactModel := new(Contact)

	communityModel := new(Community)

	createTable(userModel, contactModel, communityModel)

}
