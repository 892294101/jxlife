package api

import (
	"github.com/892294101/jxlife/server/global"
	"github.com/892294101/jxlife/server/response"
	"github.com/gin-gonic/gin"
)

type WebFileUpload struct {
}

func GetWebFileUpload() *WebFileUpload {
	return &WebFileUpload{}
}

// 图片文件上传
func (f *WebFileUpload) FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed("图片上传出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.SavePath+file.Filename)
	if err != nil {
		return
	}
	imageURL := image.AccessUrl + file.Filename
	response.Success("上传图片成功", imageURL, c)
}
