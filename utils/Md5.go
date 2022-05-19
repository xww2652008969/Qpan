package utils

import md52 "crypto/md5"

func Gedmd5(date []byte) [16]byte {
	md5 := md52.Sum(date)
	return md5
}
