package service

import (
	"errors"
	"go_chat/app/libs"
	"go_chat/app/models"
	"go_chat/origin/model"
)

type ContactService struct {
}

func (s *ContactService) AddFriend(userId, dstId int64) (err error) {

	var (
		contact *model.Contact
	)

	if userId == dstId {
		err = libs.NewReportError(errors.New("can't add yourself"))
		return
	}

	contactModel := new(model.Contact)

	if contact, err = contactModel.GetContactById(userId, dstId, model.CONCAT_CATE_USER); err != nil {
		err = libs.NewReportError(err)
		return
	}

	if contact.Id > 0 {
		err = libs.NewReportError(errors.New("you have added this user"))
		return
	}

	return

}

func (s *ContactService) SearchFriend(userId int64) (userList *[]*models.UserModel, err error) {

	contactModel := new(model.Contact)



	return
}
