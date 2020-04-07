package controllers

import (
	"encoding/json"
	"fmt"

	"oa/models"
)

// TrainingController 教育培训
type TrainingController struct {
	BaseController
}

// AddTrainingRequest 添加请求体
type AddTrainingRequest struct {
	DataName  string `json:"data_name"`
	Publisher string `json:"publisher"`
	Date      string `json:"date"`
	Content   string `json:"content"`
	Users     string `json:"users"`
}

// AddTraining 添加方法
//@Title 添加培训
//@Description 添加培训
//@Param req formData string false "添加请求"
//@Param image formData file false "图片"
//@Param file formData file false "文件"
//@Success 200
//@Failure 403
//@router /addtraining [post]
func (t *TrainingController) AddTraining() {
	req := &AddTrainingRequest{}
	var imgPath, filePath string
	str := t.GetString("req")
	fmt.Printf("%v \n", str)
	data := []byte(str)
	json.Unmarshal(data, &req)

	// 获取图像文件
	img, imgHead, _ := t.GetFile("image")
	if img != nil {
		defer img.Close()
		// 上传图片路径 "./upload/imgs/"
		// 上传图片文件
		t.SaveToFile("image", "./upload/training/images/"+imgHead.Filename)
		imgPath = "./upload/training/imgs/" + imgHead.Filename
	} else {
		imgPath = ""
	}

	//获取上传文件
	file, fileHead, _ := t.GetFile("file")
	if file != nil {
		defer file.Close()
		//上传文件路径 "./upload/file/"
		//上传附件文件
		//fmt.Println("附件路径" + uploadFilePath + filehead.Filename)
		t.SaveToFile("file", "./upload/training/files/"+fileHead.Filename)
		filePath = "./upload/training/file/" + fileHead.Filename
	} else {
		filePath = ""
	}

	err := models.CreateTraining(&models.Training{
		DataName:  req.DataName,
		Publisher: req.Publisher,
		Date:      req.Date,
		Content:   req.Content,
		Users:     req.Users,
		Image:     imgPath,
		File:      filePath,
	})

	if err != nil {
		t.Response(err)
	}
}

// GetAllTraining 获取方法
//@Title 获取所有资料
//@Description 获取所有培训资料
//@Success 200
//@Failure 403
//@router /getalldata [get]
func (t *TrainingController) GetAllTraining() {
	res := models.FetchAllTraining()
	t.Response(res)
}

// DeleteTraining 删除方法
//@Title 删除资料
//@Description 删除培训资料
//@Param id path int false "所删除资料的id"
//@Success 200
//@Failure 403
//@router /:id [delete]
func (t *TrainingController) DeleteTraining() {
	id := t.GetPathInt("id")
	err := models.DeleteTrain(id)
	if err != nil {
		t.ResponseError(err)
	}
}
