package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-blog/config"
	"go-blog/global"
	"strings"
	"time"
)

// GetToken 返回七牛的客户端上传凭证
func GetToken(q config.QiNiu) (token string, mac *qbox.Mac) {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac = qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken, mac
}

// GetConfig 返回七牛的上传配置
func GetConfig(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	//是否使用https域名
	cfg.UseHTTPS = global.Config.QiNiu.UseHttps
	//是否使用CDN上传加速
	cfg.UseCdnDomains = global.Config.QiNiu.UseCdnDomains
	return cfg
}

// UploadImage 上传图片到七牛
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启用七牛云上传")
	}

	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置AccessKey和SecretKey")
	}
	if float64(len(data)/(1024*1024)) > q.Size {
		return "", errors.New("文件超过设定的大小")
	}
	//鉴权 请参考文档：https://developer.qiniu.com/kodo/1238/go#5
	upToken, _ := GetToken(q)
	cfg := GetConfig(q)
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	datalen := int64(len(data))

	//获取当前的时间
	now := time.Now().Format("200601022150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), datalen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", q.CDN, ret.Key), nil
}

// RemoveImage 删除七牛的图片
func RemoveImage(filePath string, prefix string) (ok bool, err error) {
	q := global.Config.QiNiu
	cfg := GetConfig(q)
	_, mac := GetToken(q)
	nameList := strings.Split(filePath, "/")
	key := nameList[len(nameList)-1]
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err = bucketManager.Delete(q.Bucket, fmt.Sprintf("%s/%s", prefix, key))
	if err != nil {
		return false, err
	}
	return true, nil
}
