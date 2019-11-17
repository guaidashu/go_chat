package model

import (
	"github.com/go-xorm/xorm"
	"go_chat/app/libs"
	"time"
)

const (
	SEX_WOMEN  = "W"
	SEX_MAN    = "M"
	SEX_UNKONW = "U"
)

type UserModel struct {
	Id       int64     `xorm:"pk autoincr bigint(20)"`                       // 用户的id
	Mobile   string    `json:"mobile" xorm:"varchar(255)" form:"mobile"`     // 手机号码
	Passwd   string    `json:"passwd" xorm:"varchar(255)" form:"passwd"`     // 用户密码 = f(plainpwd + salt), MD5
	Avatar   string    `json:"avatar" xorm:"varchar(255)" form:"avatar"`     // 头像
	Sex      string    `json:"sex" xorm:"varchar(2)" form:"sex"`             //
	Nickname string    `json:"nickname" xorm:"varchar(255)" form:"nickname"` //
	Salt     string    `json:"salt" xorm:"varchar(10)" form:"salt"`          // 加密随机数
	Online   int       `json:"online" xorm:"int(10)" form:"online"`          //
	Token    string    `json:"token" xorm:"varchar(255)" form:"token"`       // /chat?id=1&token=x
	Memo     string    `json:"memo" xorm:"varchar(140)" form:"memo"`         //
	Status   int       `xorm:"int(1) default 1"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func (u *UserModel) GetDB() *xorm.Session {
	return DbEngine.Table(u.TableName())
}

func (u *UserModel) GetQueryDB() *xorm.Session {
	return DbEngine.Table(u.TableName()).Where("status=?", 1)
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) IsExists() (bool, error) {
	return DbEngine.IsTableExist(u.TableName())
}

func (u *UserModel) CreateTable() (err error) {

	// 自动建表
	err = DbEngine.Sync2(u)

	if err != nil {
		err = libs.NewReportError(err)
	}

	return

}

func (u *UserModel) GetUser(mobile string) (*UserModel, error) {

	var (
		user UserModel
		err  error
	)

	if _, err = u.GetQueryDB().Where("mobile=?", mobile).Get(&user); err != nil {
		err = libs.NewReportError(err)
	}

	return &user, err

}

func (u *UserModel) Insert() (id int64, err error) {

	db := u.GetDB()

	u.Status = 1

	id, err = db.InsertOne(u)

	if err != nil {
		err = libs.NewReportError(err)
	}

	return

}
