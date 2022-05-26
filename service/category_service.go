package service

import (
	"cms_project/model"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

//食品种类服务
//接口定义
type CategoryService interface {
	//查询商铺对应的食种类列表
	GetCategoryByShopId(shopId int) []*model.FoodCategory
	//获取所有食品种类供添加食品时进行添加
	GetAllCategory() []*model.FoodCategory
	//添加食品种类
	AddCategory(category model.FoodCategory) bool
	//添加食品
	AddFood(food model.Food) bool
	//添加商铺
	AddShop(shop model.Shop) bool
	//删除商铺
	DeleteShop(id int) bool
	//删除食品
	DeleteFood(id int) bool
}

type categoryService struct {
	Engine *xorm.Engine
}

//获取服务实体
func GetNewCategoryService(engine *xorm.Engine) CategoryService {
	return categoryService{
		Engine: engine,
	}
}

//查询商铺对应的食种类列表
func (cs categoryService) GetCategoryByShopId(shopId int) []*model.FoodCategory {
	cateList := []*model.FoodCategory{}
	cs.Engine.Where("restaurant_id = ?", shopId).Find(&cateList)
	return cateList
}

//获取所有食品种类供添加食品时进行添加
func (cs categoryService) GetAllCategory() []*model.FoodCategory {
	cateList := []*model.FoodCategory{}
	cs.Engine.Where("del_flag = 0").Find(&cateList)
	return cateList
}

//添加食品种类
func (cs categoryService) AddCategory(category model.FoodCategory) bool {
	_, err := cs.Engine.Insert(&category)
	if err != nil {
		iris.New().Logger().Error(err.Error())
	}
	return err == nil
}

//添加食品
func (cs categoryService) AddFood(food model.Food) bool {
	_, err := cs.Engine.Insert(&food)
	if err != nil {
		iris.New().Logger().Error(err.Error())
	}
	return err == nil
}

//添加商铺
func (cs categoryService) AddShop(shop model.Shop) bool {
	_, err := cs.Engine.Insert(&shop)
	if err != nil {
		iris.New().Logger().Error(err.Error())
	}
	return err == nil
}

//删除商铺
func (cs categoryService) DeleteShop(id int) bool {
	shop := model.Shop{Dele: 1}
	_, err := cs.Engine.Where("item_id=?", id).Cols("dele").Update(&shop)
	if err != nil {
		iris.New().Logger().Error(err.Error())
	}
	return err == nil
}

//删除食品
func (cs categoryService) DeleteFood(id int) bool {
	food := model.Food{DelFlag: 1}
	_, err := cs.Engine.Where("item_id = ?", id).Cols("del_flag").Update(&food)
	if err != nil {
		iris.New().Logger().Error(err.Error())
	}
	return err == nil
}
