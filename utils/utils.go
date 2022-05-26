package utils

import "time"

//请求状态码
const (
	RECODE_OK      = 1  //请求成功
	RECODE_FAIL    = 0  //请求失败
	RECODE_UNLOGIN = -1 //未登录
)

//业务逻辑状态码
const (
	RESPMSG_OK   = "1"
	RESPMSG_FAIL = "0"

	//管理员
	RESPMSG_SUCCESSLOGIN    = "SUCCESS_LOGIN"
	RESPMSG_FAILURELOGIN    = "FAILURE_LOGIN"
	RESPMSG_SUCCESSSESSION  = "SUCCESS_SESSION"
	RESPMSG_ERRORSESSION    = "ERROR_SESSION"
	RESPMSG_SIGNOUT         = "SESSION_SIGNOUT"
	RESPMSG_HASNOACCESS     = "HAS_NO_ACCESS"
	RESPMSG_ERRORANMINCOUNT = "ERROR_ADMINCOUNT"

	//用户
	RESPMSG_ERROR_USERLIST = "ERROR_USERS"
	RESPMSG_ERROR_USERINFO = "ERROR_USERINFO"

	//订单
	RESPMSG_ERROR_ORDERINFO = "ERROR_ORDERINFO"

	//商家
	RESPMSG_ERROR_RESTLIST       = "ERROR_RESTAURANTS"
	RESPMSG_SUCCESS_ADDREST      = "ADD_RESTUANT_SUCCESS"
	RESPMSG_FAIL_ADDREST         = "ADD_RESTUANT_FAIL"
	RESPMSG_ERROR_RESTAURANTINFO = "ERROR_RESTAURANTINFO"
	RESPMSG_SUCCESS_DELETESHOP   = "SUCCESS_DELETESHOP"
	RESPMSG_ERROR_SEARCHADDRESS  = "ERROR_SEARCHADDRESS"

	//食品种类
	RESPMSG_SUCCESS_CATEGORYADD = "SUCCESS_ADDCATEGORY"
	RESPMSG_ERROR_CATEGORYADD   = "ERROR_ADDCATEGORY"
	RESPMSG_ERROR_CATEGORIES    = "ERROR_CATEGORIES"

	//食品
	RESPMSG_ERROR_FOODLIST   = "ERROR_FOODS"
	RESPMSG_ERROR_FOODADD    = "ERROR_ADDFOOD"
	RESPMSG_SUCCESS_FOODADD  = "SUCCESS_ADDFOOD"
	RESPMSG_ERROR_FOODDELE   = "ERROR_DELEFOOD"
	RESPMSG_SUCCESS_FOODDELE = "SUCCESS_DELEFOOD"

	//文件操作
	RESPMSG_ERROR_PICTUREADD  = "ERROR_PICTUREADD"
	RESPMSG_ERROR_PICTURETYPE = "ERROR_PICTURETYPE"
	RESPMSG_ERROR_PICTURESIZE = "ERROR_PICTURESIZE"

	//未登录
	ERROR_UNLOGIN   = "ERROR_UNLOGIN"
	ERROR_UNKNOWERR = "8000"
)

var reCodeText = map[string]string{
	RESPMSG_OK:    "成功",
	RESPMSG_FAIL:  "失败",
	ERROR_UNLOGIN: "未登录，没有权限，请先登录",

	//管理员
	RESPMSG_SUCCESSLOGIN:    "管理员登录成功",
	RESPMSG_FAILURELOGIN:    "管理员登录失败",
	RESPMSG_SUCCESSSESSION:  "获取管理员信息成功",
	RESPMSG_ERRORSESSION:    "获取管理员信息失败",
	RESPMSG_SIGNOUT:         "退出登录",
	RESPMSG_HASNOACCESS:     "您的权限不足",
	RESPMSG_ERRORANMINCOUNT: "获取管理员总数失败",

	//用户
	RESPMSG_ERROR_USERLIST: "获取用户列表失败",
	RESPMSG_ERROR_USERINFO: "获取用户信息失败",

	//订单
	RESPMSG_ERROR_ORDERINFO: "获取订单信息失败",

	//商家
	RESPMSG_ERROR_RESTLIST:       "查询商家店铺失败",
	RESPMSG_SUCCESS_ADDREST:      "添加商家店铺成功",
	RESPMSG_FAIL_ADDREST:         "添加商家店铺失败",
	RESPMSG_ERROR_RESTAURANTINFO: "获取商家信息失败",
	RESPMSG_SUCCESS_DELETESHOP:   "删除商家成功",
	RESPMSG_ERROR_SEARCHADDRESS:  "搜索地址失败",

	//食品种类
	RESPMSG_SUCCESS_CATEGORYADD: "添加食品种类成功",
	RESPMSG_ERROR_CATEGORYADD:   "添加食品种类失败",
	RESPMSG_ERROR_CATEGORIES:    "获取食品种类失败",

	//食品
	RESPMSG_ERROR_FOODLIST:   "查询食品列表失败",
	RESPMSG_ERROR_FOODADD:    "添加食品失败",
	RESPMSG_SUCCESS_FOODADD:  "添加食品成功",
	RESPMSG_ERROR_FOODDELE:   "删除食品记录失败",
	RESPMSG_SUCCESS_FOODDELE: "删除食品记录成功",

	//图片操作
	RESPMSG_ERROR_PICTUREADD:  "图片上传失败",
	RESPMSG_ERROR_PICTURETYPE: "只支持PNG,JPG,JPEG格式的图片",
	RESPMSG_ERROR_PICTURESIZE: "图片尺寸太大,请保证在2M一下",

	//其他错误信息
	ERROR_UNKNOWERR: "服务器未知错误",
}

func RecondText2(code string) string {
	if str, ok := reCodeText[code]; ok {
		return str
	}
	return reCodeText[ERROR_UNKNOWERR]
}

//数据格式化处理
func FormatDatetime(time time.Time) string {
	return time.Format("2006-01-02 03:04:05")
}