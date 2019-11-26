package model

import (
	"github.com/go-xorm/xorm"
	"go_chat/app/libs"
)

// 加群 或者 加好友
const (
	CONCAT_CATE_USER     = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)

type Contact struct {
	Model   `xorm:"extends"`
	Ownerid int64  `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 记录是谁的 // 对端,10001
	Dstobj  int64  `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 对端信息
	Cate    int    `xorm:"int(11)" form:"cate" json:"cate"`          // 什么类型
	Memo    string `xorm:"varchar(120)" form:"memo" json:"memo"`     // 备注
}

func (c *Contact) GetDB() *xorm.Session {
	return DbEngine.Table(c.TableName())
}

func (c *Contact) GetNewSession() *xorm.Session {
	return DbEngine.NewSession()
}

func (c *Contact) GetQueryDB() *xorm.Session {
	return DbEngine.Table(c.TableName()).Where("status=?", 1)
}

func (c *Contact) TableName() string {
	return "contact"
}

func (c *Contact) IsExists() (bool, error) {
	return DbEngine.IsTableExist(c.TableName())
}

func (c *Contact) CreateTable() (err error) {

	// 自动建表
	err = DbEngine.Sync2(c)

	if err != nil {
		err = libs.NewReportError(err)
	}

	return

}

func (c *Contact) GetContactById(userId, dstId, mode int64) (contact *Contact, err error) {

	var (
		contactModel Contact
	)

	db := c.GetQueryDB()

	if _, err = db.Where("ownerid = ?", userId).
		And("dstid = ?", dstId).
		And("cate = ?", mode).
		Get(&contactModel); err != nil {
		err = libs.NewReportError(err)
	}

	// 得到一个新的事务
	session := c.GetNewSession()
	// 开始事务
	_ = session.Begin()

	// 存 自己的数据
	c.Ownerid = userId
	c.Dstobj = dstId
	c.Cate = int(mode)
	if _, err = session.InsertOne(c); err != nil {
		err = libs.NewReportError(err)
		_ = session.Rollback()
		return
	}

	// 存 对方的数据
	if _, err = session.InsertOne(&Contact{
		Ownerid: dstId,
		Dstobj:  userId,
		Cate:    int(mode),
	}); err != nil {
		err = libs.NewReportError(err)
		_ = session.Rollback()
		return
	}

	_ = session.Commit()

	return

}
