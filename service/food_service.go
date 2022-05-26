package service

import (
	"cms_project/model"

	"github.com/go-xorm/xorm"
)

//食品模块服务
//接口定义
type FoodService interface {
	//获取食品总数
	GetCount() (int64, error)
	//获取食品列表
	GetList(offset, limit int) []*model.Food
}

type foodService struct {
	Engine *xorm.Engine
}

//创建服务实例
func GetNewFoodService(engine *xorm.Engine) FoodService {
	return foodService{
		Engine: engine,
	}
}

//获取食品总数
func (fs foodService) GetCount() (int64, error) {
	count, err := fs.Engine.Where("del_flag = 0").Count()
	return count, err
}

//获取食品列表
func (fs foodService) GetList(offset, limit int) []*model.Food {
	foodList := []*model.Food{}
	fs.Engine.Where("del_flag = 0").Limit(limit, offset).Find(&foodList)
	return foodList
}
