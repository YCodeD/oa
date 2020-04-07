package models

type Training struct {
	ID int `json:"id"`
	// 资料名称
	DataName string `json:"data_name"`
	// 发布者
	Publisher string `json:"publisher"`
	// 发布日期
	Date string `json:"date"`
	// 内容
	Content string `json:"content"`
	// 使用者
	Users string `json:"users"`
	// 图片地址
	Image string `json:"image"`
	// 文件地址
	File string `json:"file"`
}

// 添加培训资料
func CreateTraining(ct *Training) error {
	err := db.Create(&ct).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据id获取资料
func FetchDataById(id int) Training {
	var res Training
	db.Where("id=?", id).First(&res)
	return res
}

// 获取所有资料
func FetchAllTraining() []Training {
	var res []Training
	db.Find(&res)
	return res
}

// 删除培训资料
func DeleteTrain(id int) error {
	dt := FetchDataById(id)
	err := db.Delete(&dt).Error
	if err != nil {
		return err
	}
	return nil
}
