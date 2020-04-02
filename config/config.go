package config

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/speps/go-hashids"
)

func Encrypt(salt string, minLength int, params []int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.Encode(params)
		if err == nil {
			return e
		}
	}
	return ""
}

// 解密
func Decrypt(salt string, minLength int, hash string) []int {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.DecodeWithError(hash)
		if err == nil {
			return e
		}
	}
	return []int{}
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return  hex.EncodeToString(h.Sum(nil))
}