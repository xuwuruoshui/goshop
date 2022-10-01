package comm

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
)

func Encode(passwd string)(salt, hashed string){
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}
	return  password.Encode(passwd, options)
}


func Decode(pwd,salt,hashedPwd string)bool{
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}
	return password.Verify(pwd, salt,hashedPwd, options)
}