package _struct

type Reply struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func GetReply(code int, data ...interface{}) *Reply {

	var result interface{}

	if len(data) < 1 {
		result = ""
	}

	return &Reply{
		Code: code,
		Data: result,
	}
}

func GetSuccess(data ...interface{}) *Reply {
	return GetReply(0, data)
}

func GetError(data ...interface{}) *Reply {
	return GetReply(-1, data)
}

type LoginReply struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
