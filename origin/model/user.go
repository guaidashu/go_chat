package model

const (
	SEX_WOMEN  = "W"
	SEX_MAN    = "M"
	SEX_UNKONW = "U"
)

type UserModel struct {
	Model
	Id       int64  `xorm:"pk autoincr bigint(20)"`                       // 用户的id
	Mobile   string `json:"mobile" xorm:"varchar(255)" form:"mobile"`     // 手机号码
	Passwd   string `json:"passwd" xorm:"varchar(255)" form:"passwd"`     // 用户密码 = f(plainpwd + salt), MD5
	Avatar   string `json:"avatar" xorm:"varchar(255)" form:"avatar"`     // 头像
	Sex      string `json:"sex" xorm:"varchar(2)" form:"sex"`             //
	Nickname string `json:"nickname" xorm:"varchar(255)" form:"nickname"` //
	Salt     string `json:"salt" xorm:"varchar(10)" form:"salt"`          // 加密随机数
	Online   int    `json:"online" xorm:"int(10)" form:"online"`          //
	Token    string `json:"token" xorm:"varchar(255)" form:"token"`       // /chat?id=1&token=x
	Memo     string `json:"memo" xorm:"varchar(140)" form:"memo"`         //
}
