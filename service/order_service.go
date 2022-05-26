package service

import (
	"cms_project/model"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

// 订单服务
// 标准的开发模式将每个实体提供的功能以接口标准的形式定义，
// 供控制层进行调用
type OrderService interface {
	//获取订单总数
	GetCount() (int64, error)
	//获取订单列表
	GetOrderList(offset, limit int) []model.OrderDetail
}

type orderService struct {
	Engine *xorm.Engine
}

func NewOrderService(engine *xorm.Engine) OrderService {
	return orderService{
		Engine: engine,
	}
}

func (os orderService) GetCount() (int64, error) {
	count, err := os.Engine.Count(new(model.OrderStatus))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (os orderService) GetOrderList(offest, limit int) []model.OrderDetail {
	orderlist := []model.OrderDetail{}

	//查询订单详情
	err := os.Engine.Table("user_order").
		Join("inner", "order_status", "order_status.status_id=user_order.order_status_id").
		Join("inner", "user", "user.id=user_order.user_id").
		Join("inner", "shop", "shop.shop_id=user_order.shop_id").
		Join("inner", "address", "address.address_id=user_order.address_id").
		Limit(limit, offest).
		Find(&orderlist)
	iris.New().Logger().Info(orderlist[0])
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		//return nil
	}
	return orderlist
}
