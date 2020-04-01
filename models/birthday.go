package models

import "fmt"

type Birthday struct {
	ID int `json:"id"`
	// 申请人
	Name string `json:"name"`
	// 申请日期
	Date string `json:"date"`
	// 活动地点
	Destination string `json:"destination"`
	// 审批 0-未审批  1-已通过  2-未通过
	Status int `json:"status" gorm:"default:0"`
}

// CreateBirthday 添加庆生申请
func CreateBirthday(cb *Birthday) Birthday {
	db.Create(&cb)
	return *cb
}

// UpdateBirthday 更新庆生申请
func UpdateBirthday(ub *Birthday) Birthday {
	db.Save(&ub)
	return *ub
}

// FetchBirthdayById 通过id获取庆生申请
func FetchBirthdayById(id int) *Birthday {
	res := Birthday{}
	db.Where("id=?", id).First(&res)
	fmt.Printf("%v \n", res)
	return &res
}

// FetchBirthdayByOption 通过选项获取庆生申请
func FetchBirthdayByOption(option int) []Birthday {
	res := []Birthday{}
	db.Where("status=?", option).Find(&res)
	return res
}
