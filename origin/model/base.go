package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type Model struct {
	Id      int64     `xorm:"pk autoincr bigint(20)"` // 用户的id
	Status  int       `xorm:"int(1) default 1"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type BaseModel interface {
	TableName() string
	CreateTable() (err error)
	GetQueryDB() *xorm.Session
	IsExists() (bool, error)
	GetDB() *xorm.Session
}
