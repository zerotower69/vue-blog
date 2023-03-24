package ctype

import "encoding/json"

type Role int

//如何在gorm中使用枚举

const (
	PermissionAdmin        Role = 1 // admin
	PermissionUser         Role = 2 // system user
	PermissionVisitor      Role = 3 // visitor
	PermissionDisabledUser Role = 4 // disabled user
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisabledUser:
		str = "被禁用的用户"
	//nat match any
	default:
		str = "其他"
	}
	return str
}
