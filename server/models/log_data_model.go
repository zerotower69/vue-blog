package models

type LogModel struct {
	MODEL
	IP        string    // 登录ip
	Addr      string    //登录的地址
	UserID    uint      //用户id
	UserModel UserModel //用户
	Nickname  string    //用户昵称
	Level     Level     //日志等级
}
