package controllers

import (
	"encoding/json"
	"fmt"
	"oa/models"
	"time"
)

// ExerciseController 锻炼小组控制器
type ExerciseController struct {
	BaseController
}

type CreateRequest struct {
	GroupName string `json:"groupname"`
	Count     int    `json:"count"`
}

// @Title 创建队伍
// @Description 创建队伍
// @Param body body controllers.CreateRequest "创建请求"
// @Success 200
// @Failure 403
// @router /createteam [post]
func (e *ExerciseController) CreateTeam() {
	cr := &CreateRequest{}
	json.Unmarshal(e.Ctx.Input.RequestBody, &cr)
	fmt.Printf("%v \n", cr)
	eg := models.ExerciseGroup{
		GroupName: cr.GroupName,
		Count:     cr.Count,
	}
	models.CreateTeam(&eg)
}

// @Title 加入队伍
// @Description 加入队伍请求
// @Param uid path int false "加入队伍用户的id"
// @Param tid path int false "加入队伍的id"
// @Success 200
// @Failure 403
// @router /join/:uid/:tid [put]
func (e *ExerciseController) JionTeam() {
	iu := e.GetPathInt("uid")
	it := e.GetPathInt("tid")

	models.JoinTeam(iu, it)
}

// @Title 取消组队
// @Description 取消组队
// @Param uid path int false "取消组队的用户id"
// @Param tid path int false "取消组队的队伍id"
// @Success 200
// @Failure 403
// @router /out/:uid/:tid [put]
func (e *ExerciseController) OutTeam() {
	iu := e.GetPathInt("uid")
	it := e.GetPathInt("tid")

	models.OutTeam(iu, it)
}

// @Title 获取记录
// @Description 获取打卡记录
// @Param name path string true 用户名
// @Success 200
// @Failure 403
// @router /getrecords/:name [get]
func (e *ExerciseController) GetRecords() {
	fmt.Println("GetRecords")
	username := e.Ctx.Input.Param(":name")
	fmt.Printf("%#v \n", username)
	records := models.GetRecords(username)
	e.Response(records)
}
func (e *ExerciseController) URLMapping() {
	e.Mapping("get", e.GetRecords)
}

// @Title 获取队伍列表
// @Description 获取队伍列表
// @Success 200
// @Failure 403
// @router /grouplist [get]
func (e *ExerciseController) GetGroupList() {
	var egs []*models.ExerciseGroup
	egs = models.GetAllGroups()
	fmt.Printf("%v", egs)
	e.Response(egs)
}

// ClockRequest 打卡请求
type ClockRequest struct {
	UserName     string `json:"username"`
	ExerciseItem string `json:"item"`
	BeginTime    string `json:"begintime"`
	EndTime      string `json:"endtime"`
	Date         string `json:"date"`
}

// @Title 锻炼打卡
// @Description 锻炼打卡
// @Param body body controllers.ClockRequest 打卡请求
// @Success 200
// @Failure 403
// @router /exerciseclock [post]
func (e *ExerciseController) ExerciseClock() {
	cr := &ClockRequest{}
	json.Unmarshal(e.Ctx.Input.RequestBody, cr)

	//"begintime": "2020-03-18 18:00:00",
	//"endtime": "2020-03-18 20:00:00",
	layout := "2006-01-02 15:04:05"
	beginTime, _ := time.ParseInLocation(layout, cr.BeginTime, time.Local)
	endTime, _ := time.ParseInLocation(layout, cr.EndTime, time.Local)
	duration := endTime.Sub(beginTime)

	ec := models.ExerciseClock{
		UserName:     cr.UserName,
		ExerciseItem: cr.ExerciseItem,
		BeginTime:    beginTime,
		EndTime:      endTime,
		Duration:     time.Duration(duration.Minutes()),
		Date:         cr.Date,
	}
	//fmt.Printf("%#v\n", ec)

	result := models.CreateClock(&ec)
	e.Response(result)
}
