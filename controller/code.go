package controller

//定义错误码类型
type Rescode int64

//定义错误码
const (
	CodeSuccess Rescode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeNeedLogin
)

var codeMasMap = map[Rescode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "参数错误",
	CodeUserExist:       "用户已经存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeInvalidToken:    "无效的token",
	CodeNeedLogin:       "需要登录",
}

//Msg()根据状态码返回错误信息，并判断该状态码是不是有对应消息，如果没有都返回服务器繁忙
func (c Rescode) Msg() string {
	msg, ok := codeMasMap[c]
	if !ok {
		msg = codeMasMap[CodeServerBusy]
	}
	return msg
}
