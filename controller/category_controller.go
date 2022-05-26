package controller

import (
	"cms_project/model"
	"cms_project/service"
	"cms_project/utils"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

//食品类别处理器
type CategoryController struct {
	Ctx     iris.Context
	Service service.CategoryService
}

// 添加食品种类实体
type CategoryEntity struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	RestaurantId string `json:"restaruant_id"`
}

//添加食品实体
type AddFoodEntity struct {
	Name         string   `json:"name"`          //食品名称
	Description  string   `json:"description"`   //食品描述
	ImagePath    string   `json:"image_path"`    //食品图片地址
	Activity     string   `json:"activity"`      //食品活动
	Attributes   []string `json:"attributes"`    //食品特点
	Specs        []Specs  `json:"specs"`         //食品规格
	CategoryId   int      `json:"category_id"`   //食品种类  种类id
	RestaurantId string   `json:"restaurant_id"` //哪个店铺的食品 店铺id
}

// 食品规格
type Specs struct {
	Specs      string `json:"specs"`
	PackingFee int    `json:"packing_fee"`
	Price      int    `json:"price"`
}

func (cc *CategoryController) BeforeActivation(a mvc.BeforeActivation) {
	//查询商铺对应的食种类列表
	a.Handle("GET", "/getcategory/{shopId}", "GetCategoryByShopId")
	// 获取全部的食品种类
	a.Handle("GET", "/v2/restaruant/category", "GetAllCategory")
	//添加商铺记录
	a.Handle("POST", "/addShop", "PostAddShop")
	//删除商铺
	a.Handle("DELETE", "/restaurant/{restaurant_id}", "DeleteRestaurant")
	//删除食品
	a.Handle("DELETE", "/v2/food/{food_id}", "DeleteFood")
}

//查询商铺对应的食种类列表
func (cc *CategoryController) GetCategoryByShopId() mvc.Result {
	shopIdStr := cc.Ctx.Params().Get("shopId")
	if shopIdStr == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORIES,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	shopId, _ := strconv.Atoi(shopIdStr)
	//调用服务
	cateList := cc.Service.GetCategoryByShopId(shopId)

	if len(cateList) <= 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORIES,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":        utils.RECODE_OK,
			"category_list": &cateList,
		},
	}
}

//获取全部食品种类
func (cc *CategoryController) GetAllCategory() mvc.Result {
	categorys := cc.Service.GetAllCategory()
	if len(categorys) <= 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORIES,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	return mvc.Response{
		Object: &categorys,
	}
}

/**
 * url：/shopping/addcategory
 * type：post
 * desc：添加食品种类记录
 */
func (cc *CategoryController) PostAddcategory() mvc.Result {
	//读取数据
	var data CategoryEntity
	cc.Ctx.ReadJSON(&data)
	if data.Name == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_CATEGORYADD),
			},
		}
	}
	//构建要添加的食品种类实例
	rId, _ := strconv.Atoi(data.RestaurantId)
	category := model.FoodCategory{
		CategoryName: data.Name,
		CategoryDesc: data.Description,
		RestaurantId: int64(rId),
	}

	ok := cc.Service.AddCategory(category)
	if !ok {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_CATEGORYADD),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORYADD,
			"message": utils.RecondText2(utils.RESPMSG_SUCCESS_CATEGORYADD),
		},
	}
}

/**
 * 添加商铺方法
 * url：/shopping/addShop
 * type：Post
 * desc：添加商铺数据记录
 */
func (cc *CategoryController) PostAddShop(ctx iris.Context) {
	//读取数据
	var shop model.Shop
	err := ctx.ReadJSON(&shop)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		ctx.JSON(iris.Map{
			"status":  utils.RECODE_FAIL,
			"type":    utils.RESPMSG_FAIL_ADDREST,
			"message": utils.RecondText2(utils.RESPMSG_FAIL_ADDREST),
		})
		return
	}
	//调用服务添加数据
	ok := cc.Service.AddShop(shop)
	if !ok {
		iris.New().Logger().Error(err.Error())
		ctx.JSON(iris.Map{
			"status":  utils.RECODE_FAIL,
			"type":    utils.RESPMSG_FAIL_ADDREST,
			"message": utils.RecondText2(utils.RESPMSG_FAIL_ADDREST),
		})
		return
	}
	//添加成功
	ctx.JSON(iris.Map{
		"status":     utils.RECODE_OK,
		"message":    utils.RecondText2(utils.RESPMSG_SUCCESS_ADDREST),
		"shopDetail": shop,
	})
}

/*
 * 添加食品信息方法
 * url：/shopping/addfood
 * type：Post
 * desc：添加食品数据记录
 */
func (cc *CategoryController) PostAddfood() mvc.Result {
	//接收添加食品实体
	var foodentity AddFoodEntity
	err := cc.Ctx.ReadJSON(&foodentity)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_FOODADD),
			},
		}
	}

	//构建添加到数据库的实体
	rId, _ := strconv.Atoi(foodentity.RestaurantId)
	food := model.Food{
		Name:         foodentity.Name,
		Description:  foodentity.Description,
		ImagePath:    foodentity.ImagePath,
		Activity:     foodentity.Activity,
		CategoryId:   int64(foodentity.CategoryId),
		RestaurantId: int64(rId),
		DelFlag:      0,
		Rating:       0,
	}
	//调用服务添加食品信息
	ok := cc.Service.AddFood(food)
	if !ok {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_FOODADD),
			},
		}
	}

	//添加成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"message": utils.RecondText2(utils.RESPMSG_SUCCESS_FOODADD),
		},
	}
}

//删除商铺
func (cc *CategoryController) DeleteRestaurant() mvc.Result {
	//获取id
	rtId := cc.Ctx.Params().Get("restaurant_id")
	id, err := strconv.Atoi(rtId)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_HASNOACCESS,
				"message": utils.RecondText2(utils.RESPMSG_HASNOACCESS),
			},
		}
	}
	ok := cc.Service.DeleteShop(id)
	if !ok {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_HASNOACCESS,
				"message": utils.RecondText2(utils.RESPMSG_HASNOACCESS),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELETESHOP,
			"message": utils.RecondText2(utils.RESPMSG_SUCCESS_DELETESHOP),
		},
	}
}

//删除食品
func (cc *CategoryController) DeleteFood() mvc.Result {
	//获取id
	foodId := cc.Ctx.Params().Get("food_id")
	id, err := strconv.Atoi(foodId)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODDELE,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_FOODDELE),
			},
		}
	}
	//调用服务
	ok := cc.Service.DeleteFood(id)
	if !ok {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODDELE,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_FOODDELE),
			},
		}
	}
	//删除成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_FOODDELE,
			"message": utils.RecondText2(utils.RESPMSG_SUCCESS_FOODDELE),
		},
	}
}
