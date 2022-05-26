package model

type OrderDetail struct {
	UserOrder   `xorm:"extends"`
	User        `xorm:"extends"`
	OrderStatus `xorm:"extends"`
	Shop        `xorm:"extends"`
	Address     `xorm:"extends"`
}

func (od *OrderDetail) OrderDetailResp() interface{} {
	resDesc := map[string]interface{}{
		"id":                   od.UserOrder.Id,
		"total_amount":         od.UserOrder.SumMoney,
		"user_name":            od.User.UserName,          //用户名
		"status":               od.OrderStatus.StatusDesc, //订单状态
		"restaurant_id":        od.Shop.Id,                //商铺id
		"restaurant_image_url": od.Shop.ImagePath,         //商铺图片
		"restaurant_name":      od.Shop.Name,              //商铺名
		"formatted_created_at": od.Time,                   //创建时间
		"status_code":          0,
		"address_id":           od.Address.AddressId, //订单地址
	}

	statusDesc := map[string]interface{}{
		"color":     "f60",
		"sub_title": "15分钟内支付",
		"title":     od.OrderStatus.StatusDesc,
	}

	resDesc["status_bar"] = statusDesc

	return resDesc
}
