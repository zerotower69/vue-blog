package flag

import sys_flag "flag"

type Option struct {
	Version bool
	DB      bool
}

// Parse the command line arguments
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

// IsWebStop stop the web project?
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigraOptions(option)
	}
}
