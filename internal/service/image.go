package service

import (
	"ai-meeting-server/internal/biz"
	"ai-meeting-server/internal/conf"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ImageService struct {
	iu   *biz.ImageUsecase
	conf *conf.Data
	log  *log.Helper
}

func NewImageService(iu *biz.ImageUsecase, conf *conf.Data, logger log.Logger) *ImageService {
	return &ImageService{
		iu:   iu,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (s *ImageService) UploadFile(ctx http.Context) error {
	req := ctx.Request()

	file, header, err := req.FormFile("file")
	if err != nil {
		return ctx.JSON(200, Resp(400, "file not found", nil))
	}
	defer file.Close()
	// 检查文件大小
	if header.Size > (10 << 20) { // 限制文件大小为10MB
		s.log.Error("file size exceeds the limit")
		return ctx.JSON(200, Resp(400, "file size exceeds the limit", nil))
	}

	// 检查文件类型
	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" { // 限制文件类型为jpg和png
		s.log.Error("unsupported file type")
		return ctx.JSON(200, Resp(400, "unsupported file type", nil))
	}
	name := strconv.Itoa(int(time.Now().UnixMicro())) + ext
	// 创建一个新的文件
	// dst, err: = os.(header.Filename)
	filePath := ""
	imgPath := ""
	if s.conf.AssetsPath != "" { // 不为空，表示使用本次存储
		imgPath = "/image/" + name
		filePath = s.conf.AssetsPath + imgPath
		// 判断目录是否存在
		dirPath := s.conf.AssetsPath + "/image"
		_, err = os.Stat(dirPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				s.log.Error("dir not exists")
				return ctx.JSON(200, Resp(500, "dir not exists", nil))
			}
		}
	} else {
		filePath = name
	}
	dst, err := os.Create(filePath)
	if err != nil {
		s.log.Error(err.Error())
		return ctx.JSON(200, Resp(500, err.Error(), nil))
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件中
	if _, err := io.Copy(dst, file); err != nil {
		s.log.Error(err.Error())
		return ctx.JSON(200, Resp(500, err.Error(), nil))
	}
	if s.conf.AssetsPath == "" { // 非本地存储，使用对象存储
		imgPath = s.iu.UploadImage(dst, name)
		if err := os.Remove(name); err != nil {
			s.log.Error(err.Error())
		}
	}
	// 使用本地存储文件
	s.log.Debugf("imgPath: %s", imgPath)
	return ctx.JSON(200, Resp(200, "success", map[string]string{
		"imageUrl": imgPath,
	}))
}

func Resp(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
