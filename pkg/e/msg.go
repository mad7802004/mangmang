package e

var ZhCnMsgFlags = map[int]string{
	SUCCESS:                  "成功",
	FAIL:                     "失败",
	NewFailed:                "新建失败",
	UpdateFailed:             "更新失败",
	FailedToDelete:           "删除失败",
	AcquisitionFailed:        "获取失败",
	InvalidParameter:         "参数无效",
	InvalidAuthorization:     "授权无效,请登录",
	VerificationCodeError:    "验证码错误",
	MobileNumberError:        "电话号码错误",
	FrequentOperation:        "请勿频繁操作",
	InconsistentPassword:     "密码不一致",
	PhoneNumberIsRegistered:  "手机号码被注册",
	AccountOrPassWordErr:     "账户或密码错误",
	OldPasswordError:         "旧密码错误",
	NoResourcesFound:         "没有找到资源",
	AccountDoesNotExist:      "账户不存在",
	BusinessCardDoesNotExist: "名片不存在",
	ImageByteIsTooLarge:      "图片最大不超过40KB",
	ImageFormatIsIncorrect:   "图片格式不正确",
	RoleDoesNotExist:         "角色不存在",
	ProjectDoesNotExist:      "项目不存在",
	ProjectUserExist:         "用户已存在该项目中，请勿重复添加",
	ProjectUserDoesNotExist:  "项目用户不存在",
	FatherTaskDoesNotExist:   "父级任务不存在",
	TaskDoesNotExit:          "任务不存在",
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
