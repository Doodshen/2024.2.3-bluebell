package controller


import ()

//定义错误码类型 
type Rescode int64


//定义错误码
const(
	CodeSuccess Rescode = 1000+iota
	CodeInvalidParam 
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMasMap = map[]