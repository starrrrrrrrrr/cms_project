package controller

import (
	"cms_project/service"
	"cms_project/utils"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

//订单处理器
type OrderController struct {
	//iris处理请求上下文
	Ctx iris.Context
	//订单模块功能实体
	Service service.OrderService
	//session实例
	//session sessions.Session
}

//订单详情
func (oc *OrderController) Get() mvc.Result {

	offset, _ := strconv.Atoi(oc.Ctx.FormValue("offset"))
	limit, err := strconv.Atoi(oc.Ctx.FormValue("limit"))

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_ORDERINFO),
			},
		}
	}

	if offset < 0 {
		offset = 0
	}
	if limit > MAXLIMIT {
		limit = MAXLIMIT
	}

	orderlist := oc.Service.GetOrderList(offset, limit)

	if orderlist == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_ORDERINFO),
			},
		}
	}

	var list []interface{}

	for _, v := range orderlist {
		list = append(list, v.OrderDetailResp())
	}

	return mvc.Response{
		Object: &list,
	}
}

//订单总数
func (oc *OrderController) GetCount() mvc.Result {
	count, err := oc.Service.GetCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_ORDERINFO),
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
