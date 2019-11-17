package model

import (
	"github.com/go-xorm/xorm"
)

type BaseModel interface {
	TableName() string
	CreateTable() (err error)
	IsExists() (bool, error)
	GetDB() *xorm.Session
}
