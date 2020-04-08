package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

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
		// 限定图片格式
		ext := path.Ext(imgHead.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			t.ResponseError("上传图片格式错误，需上传jpg，jpeg，png格式")
		}
		// 上传图片路径
		rootPath := "./upload/training/images/"
		// 判断路径是否存在
		exist := t.IsExist(rootPath)
		if exist == false {
			err := os.MkdirAll(rootPath, os.ModePerm)
			if err != nil {
				t.ResponseError(err)
			}
		}
		// 重命名
		imgName := time.Now().Format("2006-01-02-15-04-05") + ext
		// 上传图片文件
		t.SaveToFile("image", "./upload/training/images/"+imgName)
		imgPath = "./upload/training/images/" + imgName
	} else {
		imgPath = ""
	}

	//获取上传文件
	file, fileHead, _ := t.GetFile("file")
	if file != nil {
		defer file.Close()
		// 上传文件路径
		rootPath := "./upload/training/files/"
		// 判断路径是否存在
		exist := t.IsExist(rootPath)

		if exist == false {
			err := os.MkdirAll(rootPath, os.ModePerm)
			if err != nil {
				t.ResponseError(err)
			}
		}
		// 重命名
		fileName := time.Now().Format("2006-01-02-15-04-05") + path.Ext(fileHead.Filename)

		t.SaveToFile("file", "./upload/training/files/"+fileName)
		filePath = "./upload/training/files/" + fileName
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
