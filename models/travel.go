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
	Status int `json:"status" gorm:"default:0"`
}

// CreateTravel 添加出游申请
func CreateTravel(ct *Travel) Travel {
	db.Create(&ct)
	return *ct
}

// FetchTravelById 通过id获取出游申请记录
func FetchTravelById(id int) Travel{
	var res Travel
	db.Where("id=?", id).First(&res)
	return res
}

// UpdateTravel 更新出游申请
func UpdateTravel(ut *Travel) error {
	err := db.Save(&ut).Error
	if err != nil {
		return err
	}
	return nil
}

//FetchTravelByOption 通过选项获取出游申请
func FetchTravelByOption(option int) []Travel {
	res := []Travel{}
	db.Where("status=?", option).Find(&res)
	return res
}