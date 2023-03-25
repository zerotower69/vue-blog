package models

// MessageModel record messages
type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"`    //发送人ID
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`    // 发送人
	SendUserNickname string    `gorm:"size:42" json:"send_user_nickname"` // 发送人昵称
	SendUserAvatar   string    `json:"send_user_avatar"`                  //发送人头像

	ReceiveUserID       uint      `gorm:"primaryKey" json:"receive_user_id"`    //接收人ID
	ReceiveUserModel    UserModel `gorm:"foreignKey:ReceiveUserID" json:"-"`    //接收人
	ReceiveUserNickname string    `gorm:"size:42" json:"receive_user_nickname"` //接收人昵称
	ReceiveUserAvatar   string    `json:"receive_user_avatar"`                  // 接收人头像
	IsRead              bool      `gorm:"default:false" json:"is_read"`         //接收方是否查阅
	Content             string    `json:"content"`                              // 消息内容
}
