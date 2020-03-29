package controllers

import (
	"encoding/json"
	"fmt"
	"oa/models"
	"strconv"
	"time"
)

// Announcement 公告控制器
type AnnouncementController struct {
	BaseController
}

// 添加公告请求
type AddRequest struct {
	Title     string `json:"title" validate:"required"`
	Publisher string `json:"publisher"`
	Content   string `json:"content" validate:"required"`
}

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
		// 上传图片路径 "./upload/imgs/"
		// 上传图片文件
		a.SaveToFile("image", "./upload/announcement/images/"+imgHead.Filename)
		imgPath = "./upload/announcement/imgs/" + imgHead.Filename
	} else {
		imgPath = ""
	}

	//获取上传文件
	file, fileHead, err := a.GetFile("file")
	if file != nil {
		defer file.Close()
		//上传文件路径 "./upload/file/"
		//上传附件文件
		//fmt.Println("附件路径" + uploadFilePath + filehead.Filename)
		a.SaveToFile("file", "./upload/announcement/files/"+fileHead.Filename)
		filePath = "./upload/announcement/file/" + fileHead.Filename
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

// @Title 获取未读读公告
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
