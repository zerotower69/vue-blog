package v1

import (
	"go-blog/api/v1/images"
	"go-blog/api/v1/settings"
)

type ApiGroup struct {
	SettingsApi settings.SettingsApi
	ImagesApi   images.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
