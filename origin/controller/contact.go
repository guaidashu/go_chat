package controller

import (
	"go_chat/app/libs"
	"go_chat/origin/args"
	"go_chat/origin/service"
	_struct "go_chat/origin/struct"
	"go_chat/origin/utils"
	"net/http"
)

var contactService = &service.ContactService{}

func AddFriend(w http.ResponseWriter, req *http.Request) {

	var (
		arg args.ContactArg
		err error
	)

	if err = utils.Bind(req, &arg); err != nil {
		err = libs.NewReportError(err)
		_struct.WriteError(w, err)
		return
	}

	if err = contactService.AddFriend(arg.Userid, arg.Dstid); err != nil {
		err = libs.NewReportError(err)
		_struct.WriteError(w, err)
		return
	}

	_struct.WriteSuccess(w, "ok")

}
