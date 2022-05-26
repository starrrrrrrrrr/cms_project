package controller

import (
	"cms_project/model"
	"cms_project/service"
	"cms_project/utils"
	"encoding/json"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminController struct {
	//iris框架请求上下文
	Ctx iris.Context
	//admin功能实体
	Service service.AdminService
	//session对象
	Session *sessions.Session
}

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// 管理员登录功能
// 接口： /admin/login
func (ac *AdminController) PostLogin() mvc.Result {
	iris.New().Logger().Info("login")
	//创建admin实例
	adminLogin := AdminLogin{}
	//读取adminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	if adminLogin.Password == "" || adminLogin.UserName == "" {
		return mvc.Response{
			Object: iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_FAILURELOGIN,
				"message": utils.RecondText2(utils.RESPMSG_FAILURELOGIN),
			},
		}
	}

	admin, ok := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)

	if !ok {
		return mvc.Response{
			Object: iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_FAILURELOGIN,
				"message": utils.RecondText2(utils.RESPMSG_FAILURELOGIN),
			},
		}
	}

	//admin存在，将信息存入session中
	adminByte, _ := json.Marshal(&admin)
	ac.Session.Set(ADMIN, adminByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"sucess":  "登录成功",
			"message": utils.RESPMSG_SUCCESSLOGIN,
		},
	}
}

// 获取管理员信息功能
// 接口： /admin/info
func (ac *AdminController) GetInfo() mvc.Result {
	iris.New().Logger().Info("admin info")
	adminByte := ac.Session.Get(ADMIN)
	//判断session是否为空,为空用户未登录
	if adminByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.ERROR_UNLOGIN,
				"message": utils.RecondText2(utils.ERROR_UNLOGIN),
			},
		}
	}
	//创建admin实例
	admin := model.Admin{}
	//从中读取信息
	err := json.Unmarshal(adminByte.([]byte), &admin)

	//解析失败
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.ERROR_UNLOGIN,
				"message": utils.RecondText2(utils.ERROR_UNLOGIN),
			},
		}
	}

	//解析成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data":   admin.AdminToRespDesc(),
		},
	}
}

// 获取管理员总数功能
// 接口： /admin/count
func (ac *AdminController) GetCount() mvc.Result {
	//获取管理员总数
	count, err := ac.Service.GetAdminCount()
	//获取出错
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERRORANMINCOUNT,
				"message": utils.RecondText2(utils.RESPMSG_ERRORANMINCOUNT),
			},
		}
	}

	//获取成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}

}

// 管理员退出登录功能
// 接口： /admin/signout
func (ac *AdminController) GetSignout() mvc.Result {
	//删除session
	ac.Session.Delete(ADMIN)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"message": utils.RecondText2(utils.RESPMSG_SIGNOUT),
		},
	}
}
