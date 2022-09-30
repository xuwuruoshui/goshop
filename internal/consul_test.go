package internal

import (
	"github.com/google/uuid"
	"testing"
)

func TestReg(t *testing.T) {
	randUUID := uuid.New().String()
	err := Reg(AppConf.AccountWeb.Host,
		AppConf.AccountWeb.SrvName,
		randUUID,
		AppConf.AccountWeb.Port,
		AppConf.AccountWeb.Tags)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("注册成功")
}


func TestGetServerList(t *testing.T) {
	GetServerList()
}

func TestFilterService(t *testing.T) {
	FilterService()
}