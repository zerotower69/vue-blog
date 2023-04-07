package images

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/ctype"
	"go-blog/models/res"
	"go-blog/plugins/qiniu"
	"go-blog/utils"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string  `json:"file_name"`  //文件名
	Size      float64 `json:"size"`       //文件的大小
	IsSuccess bool    `json:"is_success"` //是否上传成功
	Msg       string  `json:"msg"`        //响应消息
}

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
	var fileResList []FileUploadResponse

	for _, file := range fileList {
		//先判断图片是否合法
		nameList := strings.Split(file.Filename, ".")
		//fileName := nameList[0]
		//图片的文件后缀名
		suffix := strings.ToLower(nameList[len(nameList)-1])
		//大小，以M为判断单位
		size := float64(file.Size) / float64(1024*1024)
		//global中设置了合法的图片后缀白名单
		if !utils.InList(suffix, global.WhiteImageList) {
			fileResList = append(fileResList, FileUploadResponse{
				FileName:  file.Filename,
				Size:      size,
				IsSuccess: false,
				Msg:       "上传的图片不合法,该文件后缀为:" + suffix,
			})
			continue
		}
		//拼接文件的路径
		filepath := path.Join(basePath, file.Filename)
		//判断大小（M），大于则不上传
		if size > float64(global.Config.Upload.Size) {
			fileResList = append(fileResList, FileUploadResponse{
				FileName:  file.Filename,
				Size:      size,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过%dMB,当前文件大小为：%.2f", global.Config.Upload.Size, size),
			})
			continue
		}
		//打开图片，读取内容
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		//计算图片hash
		imageHash := utils.Md5(byteData)
		//去数据库查看图片是否存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		//找到图片则不再重新上传
		if err == nil {
			//找到了
			fileResList = append(fileResList, FileUploadResponse{
				FileName:  file.Filename,
				Size:      size,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}
		var res = FileUploadResponse{
			FileName: file.Filename,
			Size:     size,
		}
		fmt.Println(global.Config.QiNiu)
		//图片上传到七牛
		if global.Config.QiNiu.Enable {
			filepath, err = qiniu.UploadImage(byteData, file.Filename, "blog")
			if err != nil {
				global.Log.Error(err)
				res.IsSuccess = false
				res.Msg = "图片上传到七牛失败"
				fileResList = append(fileResList, res)
				continue
			}
			res.IsSuccess = true
			res.Msg = "图片上传到七牛成功"
			fileResList = append(fileResList, res)
			//图片入库
			global.DB.Create(&models.BannerModel{
				Path:      filepath,
				Hash:      imageHash,
				Name:      file.Filename,
				ImageType: ctype.QiNiu,
			})
			continue
		}

		err = c.SaveUploadedFile(file, filepath)

		if err != nil {
			global.Log.Error(err)
			res.IsSuccess = false
			res.Msg = "保存图片时失败，失败原因：" + err.Error()
			fileResList = append(fileResList, res)
			continue
		}
		res.IsSuccess = true
		res.Msg = "保存成功"
		fileResList = append(fileResList, res)
		//图片入库
		global.DB.Create(&models.BannerModel{
			Path:      filepath,
			Hash:      imageHash,
			Name:      file.Filename,
			ImageType: ctype.Local,
		})
	}
	res.OkWithData(fileResList, c)
}
