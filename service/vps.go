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

func (*vps) IdExist(id uint) bool {

}

func (*vps) IpExist(ip string) bool {
	return dao.Vps{Ip: ip}.Exist()
}

func (a *vps) Set(ip string, name string, port uint) error {
	t := dao.Vps{Ip: ip, Name: name, Port: port}

	if a.IpExist(ip) {
		return t.Insert()
	}
	return t.Update()
}

func (*vps) List() ([]dao.Vps, error) {
	return dao.Vps{}.Get()
}
