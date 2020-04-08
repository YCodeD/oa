package controllers

import (
	"encoding/json"
	"fmt"
	"path"

	"os"
	"strconv"
	"time"

	"oa/models"
)

// AnnouncementController 公告控制器
type AnnouncementController struct {
	BaseController
}

// AddRequest 添加公告请求
type AddRequest struct {
	Title     string `json:"title" validate:"required"`
	Publisher string `json:"publisher"`
	Content   string `json:"content" validate:"required"`
}

//{"titile":"标题","publisher":"发布者","content":"内容"}

// AddAnnouncement 添加公告方法
// @Title 添加公告
// @Description 添加公告
// @Param req formData string false "增加请求(json格式字符串)"
// @Param image formData file false "本地图片地址"
// @Param file formData file false "本地文件地址"
// @Success 200
// @Failure 403
// @router /addannouncement [post]
func (a *AnnouncementController) AddAnnouncement() {
	anreq := &AddRequest{}
	var imgPath, filePath string

	str := a.GetString("req")

	data := []byte(str)
	json.Unmarshal(data, &anreq)

	// 获取图像文件
	img, imgHead, err := a.GetFile("image")
	if img != nil {
		defer img.Close()
		// 限定文件格式
		ext := path.Ext(imgHead.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			a.ResponseError("上传图片格式错误，需上传jpg,jpeg,png格式")
		}
		// 限定文件大小可用 imgHead.Size

		// 上传图片路径
		rootPath := "./upload/announcement/images"
		// 判断路径是否存在
		exist := a.IsExist(rootPath)
		if exist == false {
			err := os.MkdirAll(rootPath, os.ModePerm)
			if err != nil {
				fmt.Println("./upload/announcement/images mkdir error")
				a.ResponseError(err)
			}
		}

		// 对文件重命名，防止文件名重复
		imgName := time.Now().Format("2006-01-02-15-04-05") + ext
		// 上传图片文件
		a.SaveToFile("image", "./upload/announcement/images/"+imgName)
		imgPath = "./upload/announcement/images/" + imgName
	} else {
		imgPath = ""
	}

	//获取上传文件
	file, fileHead, err := a.GetFile("file")
	if file != nil {
		defer file.Close()
		// 上传文件路径
		rootPath := "./upload/announcement/files"
		// 判断路径是否存在
		exist := a.IsExist(rootPath)
		if exist == false {
			err := os.MkdirAll(rootPath, os.ModePerm)
			if err != nil {
				fmt.Println("./upload/announcement/files mkdir error")
				a.ResponseError(err)
			}
		}
		// 重命名
		fileName := time.Now().Format("2006-01-02-15-04-05") + path.Ext(fileHead.Filename)
		a.SaveToFile("file", "./upload/announcement/files/"+fileName)
		filePath = "./upload/announcement/files/" + fileName
	} else {
		filePath = ""
	}

	err = models.CreateAnnouncement(&models.Announcement{
		Title:       anreq.Title,
		Publisher:   anreq.Publisher,
		PublishTime: time.Now(),
		Content:     anreq.Content,
		Image:       imgPath,
		FilePath:    filePath,
	})
	if err != nil {
		a.Response(err)
	}
}

// GetAllAnnouncement 获取公告方法
// @Title 获取所有公告
// @Description 获取所有公告
// @Success 200
// @Failure 403 body is empty
// @router /getall [get]
func (a *AnnouncementController) GetAllAnnouncement() {
	ans, err := models.FetchAllAnnouncement()
	if err != nil {
		fmt.Println("error in GetAllAnnouncement")
	}
	a.Response(ans)
}

// GetAnnouncement 获取单个公告方法
// @Title 获取单个公告
// @Description 获取单个公告
// @Param aid path int false "获取单个公告的id"
// @Success 200
// @Failure 403 body is empty
// @router /?:aid([1-9][0-9]*) [get]
func (a *AnnouncementController) GetAnnouncement() {
	i := a.GetPathInt("aid")
	an, err := models.FetchById(i)
	if err != nil {
		a.ResponseError(err)
	}
	a.Response(an)
}

// GetReadAnnouncement 获取已读公告
// @Title 获取已读公告
// @Description 获取已读公告
// @Success 200
// @Failure 403 body is empty
// @router /read [get]
func (a *AnnouncementController) GetReadAnnouncement() {
	an, err := models.FetchAnnByOption(true)
	if err != nil {
		a.ResponseError(err)
	}
	a.Response(an)
}

// GetUnreadAnnouncement 获取未读公告
// @Title 获取未读公告
// @Description 获取未读公告
// @Success 200
// @Failure 403 body is empty
// @router /unread [get]
func (a *AnnouncementController) GetUnreadAnnouncement() {
	an, err := models.FetchAnnByOption(false)
	if err != nil {
		a.ResponseError(err)
	}
	a.Response(an)
}

// UpdateRequest 更新请求
type UpdateRequest struct {
	Title     string `json:"title" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

// UpdateAnnouncement 更新公告方法
// @Title Update Announcement
// @Description 更新公告
// @Param uid path int "更新公告id"
// @Param body body controllers.UpdateRequest "更新请求体"
// @Success 200
// @Failure 403 body is empty
// @router /:uid [put]
func (a *AnnouncementController) UpdateAnnouncement() {
	r := a.Ctx.Input.Param(":uid")
	newa := &UpdateRequest{}
	json.Unmarshal(a.Ctx.Input.RequestBody, &newa) //获取请求对的json数据
	i, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println("error in UpdateAnnouncement")
	}
	an, err := models.FetchById(i) // 获取id所对应公告内容
	an.Title = newa.Title
	an.Content = newa.Content
	an.Publisher = newa.Publisher
	//an.Image = newa.Image
	//an.PublishTime = time.Now()
	nan, err := models.UpdateAnn(an)
	a.ResponseError(err)
	a.Response(nan)
}

// DeleteAnnouncement 删除公告方法
// @Title Delete
// @Description 删除公告
// @Param id path int false "删除公告id"
// @Success 200
// @Failure 403 body is empty
// @router /delete/:id [delete]
func (a *AnnouncementController) DeleteAnnouncement() {
	i := a.GetPathInt("id")
	err := models.DeleteAnn(i)
	a.ResponseError(err)
}
