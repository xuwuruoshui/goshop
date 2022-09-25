package test

import (
	"fmt"
	"goshop/account_srv/biz"
	"testing"
	"time"
)

func TestGetMd5(t *testing.T){
	t.Log(biz.GetMd5("123456"))

	// 时间混淆
	t.Log(biz.GetMd5(fmt.Sprintf("%s%d","123456",time.Now().Unix())))
	time.Sleep(time.Second)
	t.Log(biz.GetMd5(fmt.Sprintf("%s%d","123456",time.Now().Unix())))
}

