package controller

import (
	"cms_project/service"
	"cms_project/utils"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

const (
	USERMODULE  = "user_"
	ORDERMODULE = "model_"
	ADMINMODULE = "admin_"
)

//统计模块处理器
type StatisController struct {
	//iris框架处理请求上下文
	Ctx iris.Context
	//统计模块功能实体
	Service service.StatisService
	//session对象
	Session sessions.Session
}

//获取相应模块总数
func (sc StatisController) GetCount() mvc.Result {
	//读取path
	path := sc.Ctx.Path()
	//将path中的值分别读取
	pathSlice := []string{}
	if path != "" {
		pathSlice = append(pathSlice, strings.Split(path, "/")...)
	}
	//不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//获取信息
	model := pathSlice[3]
	date := pathSlice[4]
	var result int64
	switch model {
	case "user":
		userResult := sc.Session.Get(USERMODULE + date)
		if userResult != nil {
			userResult = userResult.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  userResult,
				},
			}
		} else {
			iris.New().Logger().Error(date) //时间
			result = sc.Service.GetUserDailyCount(date)
			sc.Session.Set(USERMODULE+date, result)
		}

	case "order":
		orderResult := sc.Session.Get(ORDERMODULE + date)
		if orderResult != nil {
			orderResult = orderResult.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  orderResult,
				},
			}
		} else {
			result = sc.Service.GetOrderDailyCount(date)
			sc.Session.Set(ORDERMODULE+date, result)
		}

	case "admin":
		adminResult := sc.Session.Get(ADMINMODULE + date)
		if adminResult != nil {
			adminResult = adminResult.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  adminResult,
				},
			}
		} else {
			result = sc.Service.GetAdminDailyCount(date)
			sc.Session.Set(ADMINMODULE+date, result)
		}

	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}

}
