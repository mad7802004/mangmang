package e

var ZhCnMsgFlags = map[int]string{
	SUCCESS:               "成功",
	FAIL:                  "失败",
	InvalidParameter:      "参数无效",
	InvalidAuthorization:  "授权无效,请登录",
	VerificationCodeError: "验证码错误",
	MobileNumberError:     "电话号码错误",
	FrequentOperation:     "请勿频繁操作",
	InconsistentPassword:  "密码不一致",
	PhoneNumberIsRegistered:    "手机号码被注册",
	AccountOrPassWordErr:       "账户或密码错误",
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
