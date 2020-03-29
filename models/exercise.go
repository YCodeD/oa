package models

import (
	"fmt"
	"time"
)

// ExerciseGroup 锻炼小组
type ExerciseGroup struct {
	ID        int
	GroupName string
	Count     int
}

// EGMember 小组成员
type EGMember struct {
	ID      int
	GroupID int
	Name    string
}

// ExerciseClock 锻炼打卡记录
type ExerciseClock struct {
	ID           int
	UserName     string
	ExerciseItem string
	BeginTime    time.Time
	EndTime      time.Time
	Duration     time.Duration
	Date         string
}

// CreateTeam 创建队伍
func CreateTeam(eg *ExerciseGroup) {
	err := db.Create(&eg).Error
	if err != nil {
		fmt.Printf("%v", err)
	}
}

// JionTeam 加入队伍
func JoinTeam(uid int, tid int) {
	var m EGMember
	var t ExerciseGroup

	db.Where("id=?", uid).First(&m)
	db.Where("id=?", tid).First(&t)

	count := t.Count + 1
	db.Model(&m).Update("group_id", tid)
	db.Model(&t).Update("count", count)

}

// OutTeam 加入队伍
func OutTeam(uid int, tid int) {
	var m EGMember
	var t ExerciseGroup

	db.Where("id=?", uid).First(&m)
	db.Where("id=?", tid).First(&t)

	//m.GroupID = 0
	count := t.Count - 1
	db.Model(&m).Update("group_id", 0)
	db.Model(&t).Update("count", count)
}

// GetAllGroup 获取所有队伍
func GetAllGroups() (groups []*ExerciseGroup) {
	egs := []*ExerciseGroup{}
	err := db.Find(&egs).Error
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", egs)
	return egs
}

// CreateClock 创建打卡记录
func CreateClock(ec *ExerciseClock) ExerciseClock {
	err := db.Create(ec).Error
	if err != nil {
		panic(err)
	}
	return *ec
}

// GetClockRecord 获取用户打卡记录
func GetRecords(name string) (record []ExerciseClock) {
	var records []ExerciseClock
	db.Where("user_name=?", name).Find(&records)
	return records
}
