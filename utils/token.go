package utils

import "encoding/base64"

func GenerateToken(src string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(src))
	return sEnc
}

func DecodeToken(src string) string {
	sDec, _ := base64.StdEncoding.DecodeString(src)
	return string(sDec)
}
