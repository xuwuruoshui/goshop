package internal

import "testing"

func TestReg(t *testing.T) {
	err := Reg(ViperConf.AccountWeb.Host,
		ViperConf.AccountWeb.SrvName,
		ViperConf.AccountWeb.SrvName,
		ViperConf.AccountWeb.Port,
		ViperConf.AccountWeb.Tags)
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