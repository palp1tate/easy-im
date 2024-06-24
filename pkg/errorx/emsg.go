package errorx

var codeText = map[int]string{
	ServerCommonError: "服务器异常，稍后再尝试",
	RequestParamError: "请求参数有误",
	DbError:           "数据库繁忙，稍后再尝试",
}

func ErrMsg(errcode int) string {
	if msg, ok := codeText[errcode]; ok {
		return msg
	}
	return codeText[ServerCommonError]
}
