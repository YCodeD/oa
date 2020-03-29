package models

type Travel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// 目的地
	Destination string `json:"destination"`
	//出游时间
	TravelTime string `json:"travel_time"`
	// 申请日期
	ApplyDate string `json:"apply_date"`
	// 人员名单
	StaffList string `json:"staff_list"`
	// 审批 0-未审批  1-已通过  2-未通过
	Approval int `json:"approval" gorm:"default:0"`
}

// 添加出游申请
func CreateTravel(ct *Travel) Travel {
	db.Create(&ct)
	return *ct
}
