package response

type ResCode int64

// CodeSuccess 每在这里枚举一个，就需要在codeMsgMap里面添加错误信息
const (
	CodeSuccess ResCode = 1000 + iota
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
}

func (c ResCode) Msg() string {
	if s, ok := codeMsgMap[c]; ok {
		return s
	}
	return ""
}
