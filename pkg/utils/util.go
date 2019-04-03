package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofrs/uuid"
)

func GetUUID() string {
	userId, _ := uuid.NewV4()
	return userId.String()
}

func Md5Encrypt(data string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}
