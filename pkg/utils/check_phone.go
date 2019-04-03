package utils

import (
	"github.com/mangmang/pkg/gredis"
	"regexp"
)

// 验证手机号是否正确
func CheckPhone(phone string) bool {
	if !regexp.MustCompile("^1[3-9][0-9]{9}$").MatchString(phone) {
		return false
	}
	return true
}

// 验证手机号和验证码
func CheckPhoneCode(phone, code string, delete bool) bool {

	if !CheckPhone(phone) {
		return false
	}
	sCode, err := gredis.Hget(phone, "code")
	if err != nil || string(sCode) != code {
		return false
	}
	if delete {
		_, _ = gredis.Delete(phone)
	}
	return true
}

