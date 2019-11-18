package service

import (
	"errors"
	"go_chat/app/libs"
	"go_chat/origin/model"
)

type ContactService struct {
}

func (s *ContactService) AddFriend(userId, dstId int64) (err error) {

	if userId == dstId {
		err = libs.NewReportError(errors.New("can't add yourself"))
		return
	}

	contactModel := new(model.Contact)

	return

}
