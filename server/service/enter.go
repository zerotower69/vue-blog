package service

import "go-blog/service/image"

type ServiceGroup struct {
	ImageService image.ImageService
}

var ServiceApp = new(ServiceGroup)
