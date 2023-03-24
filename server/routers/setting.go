package routers

import (
	v1 "go-vue-blog/api/v1"
)

func (router RouterGroup) SettingsRouter() {
	settingsAPi := v1.ApiGroupApp.SettingsApi

	router.GET("settings", settingsAPi.SettingsInfoView)
}
