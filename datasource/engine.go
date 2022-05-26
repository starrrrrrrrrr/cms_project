package datasource

import (
	"cms_project/config"
	"cms_project/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewMysqlEngine() *xorm.Engine {
	//获取配置信息
	conf := config.InitConfig()
	db := conf.DataBase

	dbPath := db.User + ":" + db.Pwd + "@/" + db.Database + "?charset=utf8"

	engine, err := xorm.NewEngine(db.Drive, dbPath)
	if err != nil {
		panic(err.Error())
	}

	engine.Sync2(
		new(model.Admin),
		new(model.Permission),
		new(model.Address),
		new(model.City),
		new(model.OrderStatus),
		new(model.Shop),
		new(model.User),
		new(model.UserOrder),
		new(model.Food),
		new(model.FoodCategory),
	)

	engine.ShowSQL()
	engine.SetMaxOpenConns(10)
	return engine
}
