package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mangmang/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

const (
	AccessTokenFlag  = 0 //一次性token
	RefreshTokenFlag = 1 //刷新token
)

type Claims struct {
	UserId string `json:"user_id"`
	Flag   int8   `json:"flag"`
	jwt.StandardClaims
}

func NewToken(userID string) *Claims {
	claims := &Claims{UserId: userID}
	return claims
}

// 生成token
func (c *Claims) GenerateToken(exp, flag int, ) (string, error) {

	nowTime := time.Now()
	switch flag {
	case AccessTokenFlag:
		expireTime := nowTime.Add(time.Duration(exp) * time.Hour)
		c.Flag = AccessTokenFlag
		c.ExpiresAt = expireTime.Unix()
	case RefreshTokenFlag:
		expireTime := nowTime.AddDate(0, exp, 0)
		c.ExpiresAt = expireTime.Unix()
		c.Flag = RefreshTokenFlag
	default:
		expireTime := nowTime.Add(time.Duration(exp) * time.Hour)
		c.ExpiresAt = expireTime.Unix()
		c.Flag = AccessTokenFlag
	}

	c.Issuer = "x07"
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			if time.Now().Unix() > claims.ExpiresAt {
				return nil, err
			}
			return claims, nil
		}
	}
	return nil, err
}

// 鉴权 todo:暂时只起用户验证的功能,权限未完善
func Identify(token string) (string, bool) {

	payload, err := ParseToken(token)
	if err != nil {
		return "", false
	}
	if payload.Flag != AccessTokenFlag {
		return "", false
	}

	return payload.UserId, true

}
