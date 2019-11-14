package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type Model struct {
	Status  int       `xorm:"varchar(255) default:1"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type BaseModel interface {
	TableName() string
	CreateTable() (err error)
	IsExists() (bool, error)
	GetDB() *xorm.Session
}
