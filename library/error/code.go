package e

const (
	UNKNOWN      = 0
	UNAUTHORIZED = 401
	FORBIDDEN    = 403
	NOTFOUND     = 404
	SUCCESS      = 1000
	ERROR        = 2000
)

var MSG = map[int]string{
	UNKNOWN:      "未知状态",
	UNAUTHORIZED: "请求未被授权",
	FORBIDDEN:    "请求已被拒绝",
	NOTFOUND:     "请求无法找到",
	SUCCESS:      "请求成功",
}

func GetMsg(code int) string {
	msg, ok := MSG[code]
	if !ok {
		msg = "Message"
	}
	return msg
}
