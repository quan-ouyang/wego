package util

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strings"
)

// SignMd5 add data signatrue by MD5
func SignMd5(data, secretKey string) string {
	data, err := url.QueryUnescape(data)
	if err != nil {
		return ""
	}
	data = data + "&key=" + secretKey
	m := md5.New()
	m.Write([]byte(data))
	data = hex.EncodeToString(m.Sum(nil))

	sign := strings.ToUpper(data)
	return sign
}

// CheckSignMd5 check data signatrue by MD5
func CheckSignMd5(data, secretKey, sign string) bool {
	return SignMd5(data, secretKey) == sign
}
