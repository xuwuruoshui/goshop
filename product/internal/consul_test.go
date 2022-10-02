package internal

import (
	"github.com/google/uuid"
	"testing"
)

func TestReg(t *testing.T) {
	randUUID := uuid.New().String()
	err := Reg(AppConf.ProductWeb.Host,
		AppConf.ProductWeb.SrvName,
		randUUID,
		AppConf.ProductWeb.Port,
		AppConf.ProductWeb.Tags)
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