package init

import (
	"fmt"
	"go_chat/app/libs"
	"go_chat/origin/model"
)

func init() {

	var err error

	// 自动建表
	err = model.DbEngine.Sync2(new(model.UserModel))

	if err != nil {
		libs.DebugPrint(fmt.Sprintf("%v", err.Error()))
	}

}
