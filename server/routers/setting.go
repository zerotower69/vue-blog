package routers

import (
	"go-blog/api/v1"
)

func (router RouterGroup) SettingsRouter() {
	settingsAPi := v1.ApiGroupApp.SettingsApi

	router.GET("settings", settingsAPi.SettingsInfoView)
	router.PUT("settings", settingsAPi.SettingsInfoUpdateView)
}
