package test

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
	"testing"
)

func TestGetMd5(t *testing.T){
	//t.Log(biz.GetMd5("123456"))

	// 时间混淆
	//t.Log(biz.GetMd5(fmt.Sprintf("%s%d","123456",time.Now().Unix())))
	//time.Sleep(time.Second)
	//t.Log(biz.GetMd5(fmt.Sprintf("%s%d","123456",time.Now().Unix())))

	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}
	salt, hashed := password.Encode("123456", options)
	t.Log(salt)
	t.Log(hashed)

	check := password.Verify("123456", salt, hashed, options)
	t.Log(check)

}

