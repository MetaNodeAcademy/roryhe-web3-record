package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5String(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func ValidateMD5(str, knownHash string) bool {
	computed := MD5String(str)
	// 如果你希望忽略大小写，可用 strings.EqualFold
	return computed == knownHash
}
