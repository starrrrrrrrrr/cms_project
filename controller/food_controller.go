package controller

import (
	"cms_project/service"
	"cms_project/utils"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

//食品模块处理器
type FoodController struct {
	Ctx     iris.Context
	Service service.FoodService
}

//获取所有食品总数
func (fc *FoodController) GetCount() mvc.Result {
	count, err := fc.Service.GetCount()
	if err != nil {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
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

//获取食品列表
func (fc *FoodController) Get() mvc.Result {
	//获取参数
	offset, _ := strconv.Atoi(fc.Ctx.Params().Get("offset"))
	limit, err := strconv.Atoi(fc.Ctx.Params().Get("limit"))
	if err != nil {
		offset = 0
		limit = 20
	}

	foodList := fc.Service.GetList(offset, limit)
	if len(foodList) <= 0 {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODLIST,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_FOODLIST),
			},
		}
	}

	//获取成功
	return mvc.Response{
		Object: &foodList,
	}
}
