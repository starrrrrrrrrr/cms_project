package main

import (
	"cms_project/config"
	"cms_project/controller"
	"cms_project/datasource"
	"cms_project/model"
	"cms_project/service"
	"cms_project/utils"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := NewApp()

	//app应用配置
	Configuration(app)

	//路由配置
	MvcHandle(app)

	//初始化
	conf := config.InitConfig()

	addr := ":" + conf.Port

	app.Run(
		iris.Addr(addr), //监听端口9000
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)

	//mvc架构

}

//构建app
func NewApp() *iris.Application {
	app := iris.New()

	//设置字符
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "utf-8",
	}))

	//设置图标
	app.Favicon("./static/favicons/favicon.ico")

	//设置日志级别 开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manmger/static", "./static")
	app.HandleDir("/img", "./uploads")

	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	return app

}

func Configuration(app *iris.Application) {
	//设置字符
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "utf-8",
	}))

	//错误配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errMsg": iris.StatusNotFound,
			"msg":    "not found",
			"data":   iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errMsg": iris.StatusInternalServerError,
			"msg":    "internal server error",
			"data":   iris.Map{},
		})
	})
}

func MvcHandle(app *iris.Application) {
	//启用session
	sessManger := sessions.New(sessions.Config{
		Cookie:  "seessioncookie",
		Expires: 24 * time.Hour,
	})

	//创建redis实例
	redis := datasource.NewRedis()
	//将session同步到redis
	sessManger.UseDatabase(redis)

	//创建数据库引擎
	engine := datasource.NewMysqlEngine()

	//管理员模块
	//将数据库引擎添加到AdminService
	adminService := service.NewAdminService(engine)
	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManger.Start,
	)

	admin.Handle(new(controller.AdminController))

	//用户模块
	userService := service.GetNewUserService(engine)
	user := mvc.New(app.Party("/v1/users"))
	user.Register(
		userService,
		sessManger,
	)
	user.Handle(new(controller.UserController))

	//获取用户信息
	app.Get("/v1/user/{user_name}", func(ctx iris.Context) {
		username := ctx.FormValue("user_name")
		if username == "" {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_USERINFO),
			})
		}
		var userInfo model.User
		_, err := engine.Where("user_name=?", username).Get(&userInfo)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_USERINFO),
			})
		}

		ctx.JSON(userInfo)
	})

	//统计模块
	//将数据库引擎添加到AdminService
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{mode}/{date}/"))
	statis.Register(
		statisService,
		sessManger,
	)
	statis.Handle(new(controller.StatisController))

	//订单模块
	orderService := service.NewOrderService(engine)
	order := mvc.New(app.Party("/bos/order/"))
	order.Register(
		orderService,
		sessManger,
	)
	order.Handle(new(controller.OrderController))

	//商铺模块
	shopService := service.GetNewShopService(engine)
	shop := mvc.New(app.Party("/shopping/restaurants/"))
	shop.Register(
		shopService,
	)
	shop.Handle(new(controller.ShopController))

	//食品模块
	foodService := service.GetNewFoodService(engine)
	food := mvc.New(app.Party("/shopping/v2/foods/"))
	food.Register(
		foodService,
	)
	food.Handle(new(controller.FoodController))

	// 添加食品类别
	categoryService := service.GetNewCategoryService(engine)
	category := mvc.New(app.Party("/shopping"))
	category.Register(
		categoryService,
	)
	category.Handle(new(controller.CategoryController))

	//poi地址检索
	app.Get("v1/pois/", func(ctx iris.Context) {
		//获取参数
		path := ctx.Request().URL.String()
		app.Logger().Info(path)
		//第三方搜索
		rs, err := http.Get("https://elm.cangdu.org" + path)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_SEARCHADDRESS,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_SEARCHADDRESS),
			})
			return
		}
		//读取内容
		body, err := ioutil.ReadAll(rs.Body)
		if err != nil {
			app.Logger().Error(err.Error())
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_SEARCHADDRESS,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_SEARCHADDRESS),
			})
			return
		}

		var list []*model.PoiSearch
		json.Unmarshal(body, &list)
		//返回数据
		ctx.JSON(&list)

	})

	//获取地址信息
	app.Get("v1/address/{address_id}", func(ctx iris.Context) {
		addressId := ctx.Params().Get("address_id")
		id, err := strconv.Atoi(addressId)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}

		var address model.Address
		_, err = engine.Id(id).Get(&address)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}

		//获取成功
		ctx.JSON(address)
	})

	//上传图片
	app.Post("/v1/addimg/{model}", func(ctx iris.Context) {
		//模块
		model := ctx.Params().Get("model")
		//获取文件
		file, info, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		fileName := info.Filename
		defer file.Close()
		//打开文件
		out, err := os.OpenFile("./uploads/"+model+"/"+fileName, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		defer out.Close()
		//保存文件
		io.Copy(out, file)

		ctx.JSON(iris.Map{
			"status":   utils.RECODE_OK,
			"img_path": fileName,
		})

	})

	//上传头像
	app.Post("/admin/update/avatar/{adminId}", func(ctx iris.Context) {
		//获取信息
		adminId := ctx.Params().Get("adminId")
		file, info, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		defer file.Close()
		fileName := info.Filename
		//创建文件
		of, _ := os.OpenFile("./uploads/"+fileName, os.O_CREATE|os.O_WRONLY, 0666)
		defer of.Close()
		app.Logger().Info("文件上传" + of.Name())
		//保存文件
		_, err = io.Copy(of, file)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"message": utils.RecondText2(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		//更新数据库
		intAdId, _ := strconv.Atoi(adminId)
		adminService.SaveAvatarImg(intAdId, fileName)

		ctx.JSON(iris.Map{
			"status":   utils.RECODE_OK,
			"img_path": fileName,
		})

	})
}
