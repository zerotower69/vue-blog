package images

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models/res"
	"go-blog/service"
	"go-blog/service/image"
	"io/fs"
	"os"
	"path"
)

// ImageUpload 上传多张图片
func (ImagesApi) ImageUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在该文件", c)
		return
	}
	//判断路径是否存在，不存在就创建路径
	//pathList := strings.Split(global.Config.Upload.Path,"/")
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(global.Config.Upload.Path, fs.ModePerm) //MkdirAll 节省代码量
		if err != nil {
			global.Log.Error(err)
		}
	}
	//判断的文件的大小
	var fileResList []image.FileUploadResponse
	for _, file := range fileList {
		filePath := path.Join(basePath, file.Filename)
		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			fileResList = append(fileResList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			//本地还需要保存
			err = c.SaveUploadedFile(file, filePath)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				fileResList = append(fileResList, serviceRes)
				continue
			}
		}

		fileResList = append(fileResList, serviceRes)
	}
	res.OkWithData(fileResList, c)
}
