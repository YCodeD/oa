package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

// 基础控制器
type BaseController struct {
	beego.Controller
}

const (
	// CommonError 普通
	CommonError = 1
)

// Response 响应
func (c *BaseController) Response(ps ...interface{}) {
	rsp := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{}
	switch len(ps) {
	case 3:
		rsp.Code = ps[2].(int)
		fallthrough
	case 2:
		switch ps[1].(type) {
		case string:
			rsp.Msg = ps[1].(string)
		case error:
			rsp.Msg = ps[1].(error).Error()
		}
		fallthrough
	case 1:
		rsp.Data = ps[0]
	}
	c.Data["json"] = rsp
	c.ServeJSON()
	c.StopRun()
}

// ResponseError响应错误
func (c *BaseController) ResponseError(ps ...interface{}) *BaseController {
	switch len(ps) {
	case 2:
		c.Response(nil, ps[0], ps[1])
	case 1:
		if ps[0] == nil {
			return c
		}
		c.Response(nil, ps[0], CommonError)
	case 0:
		c.Response(nil, "fail", CommonError)
	}
	return nil
}

// GetPathInt 将获取参数转换成int类型
func (c *BaseController) GetPathInt(v string) int {
	r := c.Ctx.Input.Param(":" + v)
	if r == "" {
		return 0
	}

	i, _ := strconv.Atoi(r)
	//fmt.Printf("2 %#v\n", i)
	return i
}
