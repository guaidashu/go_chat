package model

import "time"

type Model struct {
	Status  int       `xorm:"varchar(255) default:1"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
