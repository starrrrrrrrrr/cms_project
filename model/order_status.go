package model

/**
 * 订单状态结构体定义
 */
type OrderStatus struct {
	StatusId   int64  `xorm:"pk antoincr" json:"id"` //订单状态编号
	StatusDesc string `xorm:"varchar(255)"`          // 订单状态描述

}

/*
未支付
已支付
已发货
正在配送
已接收
发起退款
正在退款
取消订单
*/
