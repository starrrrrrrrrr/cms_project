package controller

import (
	"cms_project/service"
	"cms_project/utils"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

//商铺模块处理器实例
type ShopController struct {
	//iris请求上下文
	Ctx iris.Context
	//服务实体
	Service service.ShopService
	//session对象
	Session sessions.Session
}

//获取商铺总数
func (sc *ShopController) GetCount() mvc.Result {
	count, err := sc.Service.GetCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_RESTAURANTINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_RESTAURANTINFO),
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

// 获取商铺列表 请求url/shopping/resaurant
//请求类型;get
func (sc *ShopController) Get() mvc.Result {
	//获取参数
	offStr := sc.Ctx.FormValue("offset")
	limStr := sc.Ctx.FormValue("limit")
	if offStr == "" || limStr == "" {
		offStr = "0"
		limStr = "20"
	}

	offset, _ := strconv.Atoi(offStr)
	limit, err := strconv.Atoi(limStr)
	if err != nil {
		offset = 0
		limit = 20
	}

	//调用服务获取信息
	list := sc.Service.GetShopList(offset, limit)
	if len(list) <= 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_RESTAURANTINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_RESTAURANTINFO),
			},
		}
	}

	//转换成前端需要的信息
	var resList []interface{}
	for _, v := range list {
		resList = append(resList, v.ShopToRespDesc())
	}
	return mvc.Response{
		Object: resList,
	}
}
