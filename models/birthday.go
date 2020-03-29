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
	Approval int `json:"approval" gorm:"default:0"`
}

// 添加生日申请
func CreateBirthday(cb *Birthday) Birthday {
	db.Create(&cb)
	return *cb
}

// 更新生日申请
func UpdateBirthday(ub *Birthday) Birthday {
	db.Save(&ub)
	return *ub
}

// 通过id获取
func FetchBirthdayById(id int) *Birthday {
	res := Birthday{}
	db.Where("id=?", id).First(&res)
	fmt.Printf("%v \n", res)
	return &res
}
