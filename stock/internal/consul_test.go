package internal

import (
	"github.com/google/uuid"
	"testing"
)

func TestReg(t *testing.T) {
	randUUID := uuid.New().String()
	err := Reg(AppConf.StockWeb.Host,
		AppConf.StockWeb.SrvName,
		randUUID,
		AppConf.StockWeb.Port,
		AppConf.StockWeb.Tags)
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