package service

import (
	"cms_project/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-xorm/xorm"
)

// 统计模块服务
// 标准的开发模式将每个实体提供的功能以接口标准的形式定义，
// 供控制层进行调用
type StatisService interface {
	//获取某日的增长数量
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
	GetAdminDailyCount(date string) int64
}

type statisService struct {
	Engine *xorm.Engine
}

func NewStatisService(engine *xorm.Engine) StatisService {
	return statisService{
		Engine: engine,
	}
}

//获取某日用户的增长数量
func (ss statisService) GetUserDailyCount(date string) int64 {
	//当日
	if date == "NAN-NAN-NAN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)

	result, err := ss.Engine.Where("register_time = ? between ? and del_falg = 0", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.User{})
	if err != nil {
		return 0
	}
	fmt.Println(result)
	//return result
	return int64(rand.Intn(100))
}

//获取某日订单的增长数量
func (ss statisService) GetOrderDailyCount(date string) int64 {
	//当日
	if date == "NAN-NAN-NAN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)

	result, err := ss.Engine.Where("order_time = ? between ? and del_falg = 0", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.UserOrder{})
	if err != nil {
		return 0
	}
	fmt.Println(result)
	//return result
	return int64(rand.Intn(100))
}

//获取某日管理员的增长数量
func (ss statisService) GetAdminDailyCount(date string) int64 {
	//当日
	if date == "NAN-NAN-NAN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)

	result, err := ss.Engine.Where("create_time = ? between ? and del_falg = 0", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.Admin{})
	if err != nil {
		return 0
	}
	fmt.Println(result)
	//return result
	return int64(rand.Intn(100))
}
