package internal

import "testing"

func TestReg(t *testing.T) {
	err := Reg(AppConf.AccountWeb.Host,
		AppConf.AccountWeb.SrvName,
		AppConf.AccountWeb.SrvName,
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