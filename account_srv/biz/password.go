package biz

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(s string)string{
	sum := md5.Sum([]byte(s))

	return fmt.Sprintf("%x",sum)
}
