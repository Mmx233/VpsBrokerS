package service

import "github.com/Mmx233/VpsBrokerS/models/dao"

type vps struct{}

// Vps 客户端别名记录器
var Vps vps

func (*vps) NameExist(name string) (bool, string, error) {
	t := dao.Vps{Name: name}
	if e := t.Find(); e != nil {
		return false, "", e
	}

	return true, t.Name, nil
}

func (*vps) IpExist(ip string) bool {
	return dao.Vps{Ip: ip}.Exist()
}

func (a *vps) SetName(ip string, name string) error {
	t := dao.Vps{Ip: ip, Name: name}

	if a.IpExist(ip) {
		return t.Insert(nil)
	}
	return t.UpdateName()
}
