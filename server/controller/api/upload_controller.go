package api

import (
	"bbs-go/controller/base"
	"bbs-go/util/uploader"
	"io/ioutil"
	"net/http"
	"strconv"

	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"

	"bbs-go/service"
)

const uploadMaxM = 10
const uploadMaxBytes int64 = 1024 * 1024 * 1024 * uploadMaxM

type UploadController struct {
	base.BaseController
}

func (c *UploadController) Post(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	defer file.Close()

	if header.Size > uploadMaxBytes {
		c.JsonErrorMsg(ctx, "图片不能超过"+strconv.Itoa(uploadMaxM)+"M")
		return
	}

	contentType := header.Header.Get("Content-Type")
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	logrus.Info("上传文件：", header.Filename, " size:", header.Size)

	url, err := uploader.PutImage(fileBytes, contentType)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("url", url).Build())
	return
}

// vditor上传
func (c *UploadController) PostEditor(ctx *gin.Context) {
	errFiles := make([]string, 0)
	succMap := make(map[string]string)

	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"msg":  err.Message,
			"code": err.Code,
			"data": gin.H{
				"errFiles": errFiles,
				"succMap":  succMap,
			},
		})

	}

	form, _ := ctx.MultipartForm() //check
	files := form.File["file[]"]
	for _, file := range files {
		contentType := file.Header.Get("Content-Type")
		f, err := file.Open()
		if err != nil {
			logrus.Error(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}
		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			logrus.Error(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}
		url, err := uploader.PutImage(fileBytes, contentType)
		if err != nil {
			logrus.Error(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}

		succMap[file.Filename] = url
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"msg":  "",
		"code": 0,
		"data": gin.H{
			"errFiles": errFiles,
			"succMap":  succMap,
		},
	})
}

// vditor 拷贝第三方图片
func (c *UploadController) PostFetch(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"msg":  err.Message,
			"code": err.Code,
			"data": gin.H{},
		})
		return
	}

	data := make(map[string]string)
	err := simple.ReadForm(ctx, &data) //
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 0,
			"data": gin.H{},
		})
		return
	}

	url := data["url"]
	output, err := uploader.CopyImage(url)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 0,
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"msg":  "",
		"code": 0,
		"data": gin.H{
			"originalURL": url,
			"url":         output,
		},
	})
}
