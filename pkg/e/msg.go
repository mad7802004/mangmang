package e

var ZhCnMsgFlags = map[int]string{
	SUCCESS:              "成功",
	FAIL:                 "失败",
	InvalidAuthorization: "授权无效,请登录",
}

var EnMsgFlags = map[int]string{
	SUCCESS: "success",
	FAIL:    "fail",
}

// 获取不同版本的语言提示
func GetMsg(code int, lang string) string {
	switch lang {
	case "zh_Cn":
		msg, ok := ZhCnMsgFlags[code]
		if ok {
			return msg
		}
	case "en":
		msg, ok := EnMsgFlags[code]
		if ok {
			return msg
		}
	default:
		msg, ok := ZhCnMsgFlags[code]
		if ok {
			return msg
		}

	}
	return ZhCnMsgFlags[FAIL]
}
