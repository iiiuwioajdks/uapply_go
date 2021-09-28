package response

type ResCode int64

// CodeSuccess 每在这里枚举一个，就需要在codeMsgMap里面添加错误信息
const (
	CodeSuccess ResCode = 1000 + iota
	CodeParamsInvalid
	CodeParamsFalse
	CodeSystemBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeParamsInvalid: "invalid params",
	CodeParamsFalse:   "账号或密码错误",
	CodeSystemBusy:    "系统繁忙",
}

func (c ResCode) Msg() string {
	if s, ok := codeMsgMap[c]; ok {
		return s
	}
	return ""
}
