package image

import (
	"fmt"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/ctype"
	"go-blog/plugins/qiniu"
	"go-blog/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string  `json:"file_name"`  //文件名
	Size      float64 `json:"size"`       //文件的大小
	IsSuccess bool    `json:"is_success"` //是否上传成功
	Msg       string  `json:"msg"`        //响应消息
}

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	//先判断图片是否合法
	nameList := strings.Split(fileName, ".")
	//图片的文件后缀名
	suffix := strings.ToLower(nameList[len(nameList)-1])
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, fileName)
	res.FileName = filePath
	//大小，以M为判断单位
	size := float64(file.Size) / float64(1024*1024)
	res.Size = size
	//判断图片的后缀是否在白名单内
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = fmt.Sprintf("上传的图片不合法,该文件后缀为: %s", suffix)
		return
	}
	//判断大小（M），大于则不上传
	if size > float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过%dMB,当前文件大小为：%.2f", global.Config.Upload.Size, size)
		return
	}
	//打开图片，读取内容
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		return
	}
	byteData, err := io.ReadAll(fileObj)
	//计算图片hash
	imageHash := utils.Md5(byteData)
	//去数据库查看图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	fmt.Println(err)
	if err == nil {
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	//图片上传到七牛
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, file.Filename, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = fmt.Sprintf("图片上传到七牛失败,详细信息：%s", err.Error())
			return
		}
		protocol := "http://"
		if global.Config.QiNiu.UseHttps {
			protocol = "https://"
		}
		fileName = fmt.Sprintf("%s%s/%s", protocol, global.Config.QiNiu.CDN, filePath)
		res.FileName = fileName
		res.Msg = "图片上传到七牛成功"
		fileType = ctype.QiNiu
	} else {
		//图片保存到本地的逻辑由外部完成了
		fileName = filePath
	}

	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
