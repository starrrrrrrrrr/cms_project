package controller

import (
	"cms_project/service"
	"cms_project/utils"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

const MAXLIMIT = 50

//用户处理器
type UserController struct {
	//iris请求上下文
	Ctx iris.Context
	//user模块功能实体
	Service service.UserService
	//sessioin实例
	//session sessions.Session
}

//获取用户总数
// 请求类型：Get
// 请求URl:/v1/users/count
func (uc *UserController) GetCount(ctx iris.Context) mvc.Result {
	count, err := uc.Service.GetUserTotalCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_USERINFO),
			},
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}
}

// 获取用户总数
// 请求类型：Get
// 请求Url:/v1/users/list
func (uc *UserController) GetList(ctx iris.Context) mvc.Result {
	//获取参数
	offsetStr := uc.Ctx.FormValue("offset")
	limitStr := uc.Ctx.FormValue("limit")

	if offsetStr == "" || limitStr == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}
	offset, _ := strconv.Atoi(offsetStr)
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}

	//边界限制
	if offset < 0 {
		offset = 0
	}
	if limit > MAXLIMIT {
		limit = MAXLIMIT
	}

	//调用服务获取数据
	userList := uc.Service.GetUserList(offset, limit)
	if userList == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}

	//转换为前端需要的数据格式
	var resMsg []interface{}
	for _, v := range userList {
		resMsg = append(resMsg, v.UserToRespDesc())
	}

	return mvc.Response{
		Object: &resMsg,
	}
}
