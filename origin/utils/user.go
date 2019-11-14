package utils

import "github.com/guaidashu/go_helper/crypto_tool"

func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return crypto_tool.Md5(plainpwd+salt) == passwd
}

func MakePasswd(plainpwd, salt string) string {
	return crypto_tool.Md5(plainpwd + salt)
}
