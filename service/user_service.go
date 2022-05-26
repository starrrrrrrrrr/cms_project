package service

import (
	"cms_project/model"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

//用户功能服务接口定义
type UserService interface {
	//获取用户日增长统计数据
	GetUserDailyStatisCount(date string) int64
	//获取用户总数
	GetUserTotalCount() (int64, error)
	//用户列表
	GetUserList(offset, limit int) []*model.User
}

type userService struct {
	Engine *xorm.Engine
}

//获取UserService实例
func GetNewUserService(engine *xorm.Engine) UserService {
	return userService{
		Engine: engine,
	}
}

//获取用户日增长统计数据
func (us userService) GetUserDailyStatisCount(date string) int64 {
	//当日
	if date == "NAN-NAN-NAN" {
		date = time.Now().Format("2006-01-02")
	}

	start, err := time.Parse("2006-01-02", date)
	if err != nil {

		iris.New().Logger().Error(err)
		return 0
	}
	end := start.AddDate(0, 0, 1)

	result, err := us.Engine.Where("create_time between ? and ? and del_flag = 0", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")).
		Count()
	if err != nil {

		iris.New().Logger().Error(err)
		return 0
	}

	return result

}

/**
 * 请求用户列表数据
 * offset：偏移数量
 * limit：一次请求获取的数据条数
 */
func (us userService) GetUserList(offset, limit int) []*model.User {
	userList := []*model.User{}
	err := us.Engine.Where("del_flag=?", 0).Limit(limit, offset).Find(&userList)
	if err != nil {
		return nil
	}
	return userList
}

//获取用户总数
func (us userService) GetUserTotalCount() (int64, error) {
	count, err := us.Engine.Count(new(model.User))
	if err != nil {
		return 0, err
	}
	return count, nil
}
