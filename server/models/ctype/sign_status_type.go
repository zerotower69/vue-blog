package ctype

import "encoding/json"

type SignStatus int

//judge user signed from other platform!

const (
	SignQQ     SignStatus = 1 //QQ
	SignGitee  SignStatus = 2 //gitee
	SignGithub SignStatus = 3 //github
	SignBaidu  SignStatus = 4 //baidu
	SignEmail  SignStatus = 5 //email
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "gitee"
	case SignGithub:
		str = "github"
	case SignBaidu:
		str = "百度"
	case SignEmail:
		str = "邮箱"
	//nat match any
	default:
		str = "其他"
	}
	return str
}
