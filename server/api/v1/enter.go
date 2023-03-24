package v1

import "go-vue-blog/api/v1/settings"

type ApiGroup struct {
	SettingsApi settings.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
