package model

import (
	"github.com/go-xorm/xorm"
	"go_chat/app/libs"
)

const (
	COMMUNITY_CATE_COM = 0x01
)

type Community struct {
	Model   `xorm:"extends"`
	Name    string `xorm:"varchar(30)" form:"name" json:"name"`      // 名称
	Ownerid int64  `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 群主ID
	Icon    string `xorm:"varchar(250)" form:"icon" json:"icon"`     // 群logo
	Cate    int    `xorm:"int(11)" form:"cate" json:"cate"`          // como
	Memo    string `xorm:"varchar(120)" form:"memo" json:"memo"`     // 描述
}

func (c *Community) GetDB() *xorm.Session {
	return DbEngine.Table(c.TableName())
}

func (c *Community) GetQueryDB() *xorm.Session {
	return DbEngine.Table(c.TableName()).Where("status=?", 1)
}

func (c *Community) TableName() string {
	return "community"
}

func (c *Community) IsExists() (bool, error) {
	return DbEngine.IsTableExist(c.TableName())
}

func (c *Community) CreateTable() (err error) {

	// 自动建表
	err = DbEngine.Sync2(c)

	if err != nil {
		err = libs.NewReportError(err)
	}

	return

}
