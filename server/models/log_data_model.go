package models

type LogDataModel struct {
	MODEL
	IP        string    `gorm:"size:20" json:"ip"`          // 登录ip
	Addr      string    `gorm:"size:64" json:"addr"`        //登录的地址
	UserID    uint      `json:"user_id"`                    //用户id
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"` //用户
	Nickname  string    `gorm:"size:42" json:"nickname"`    //用户昵称
	Token     string    `gorm:"size:256" json:"token"`      //token
	Device    string    `gorm:"size:256" json:"device"`     //登录的设备
}
