package service

import (
	"cms_project/model"

	"github.com/go-xorm/xorm"
)

//商铺模块服务
//接口定义
type ShopService interface {
	//获取商铺总数
	GetCount() (int64, error)
	//获取商铺信息列表
	GetShopList(offset, limit int) []*model.Shop
}

type shopService struct {
	Engine *xorm.Engine
}

//获取商铺服务实例
func GetNewShopService(engine *xorm.Engine) ShopService {
	return &shopService{
		Engine: engine,
	}
}

//获取商铺总数
func (ss *shopService) GetCount() (int64, error) {
	count, err := ss.Engine.Where("dele = 0").Count(new(model.Shop))
	return count, err
}

//获取商铺信息列表
func (ss *shopService) GetShopList(offset, limit int) []*model.Shop {
	sList := []*model.Shop{}
	ss.Engine.Where("dele = 0").Limit(limit, offset).Find(&sList)
	return sList
}
