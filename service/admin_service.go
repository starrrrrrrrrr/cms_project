package service

import (
	"cms_project/model"

	"github.com/go-xorm/xorm"
)

// 管理员服务
// 标准的开发模式将每个实体提供的功能以接口标准的形式定义，
// 供控制层进行调用
type AdminService interface {
	//通过用户名和密码取出admin信息
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
	//获取管理员总数
	GetAdminCount() (int64, error)
	//更新头像
	SaveAvatarImg(adminId int, fileName string) bool
}

func NewAdminService(engine *xorm.Engine) AdminService {
	return &adminService{
		Engine: engine,
	}
}

type adminService struct {
	Engine *xorm.Engine
}

func (ac *adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	admin := model.Admin{}
	ac.Engine.Where("admin_name=?", username).And("pwd=?", password).Get(&admin)
	return admin, admin.AdminId != 0
}

func (ac *adminService) GetAdminCount() (int64, error) {
	count, err := ac.Engine.Count(new(model.Admin))
	if err != nil {
		return 0, err
	}
	return count, nil
}

//保存头像
func (ac *adminService) SaveAvatarImg(adminId int, fileName string) bool {
	admin := model.Admin{AdminId: int64(adminId), Avatar: fileName}
	_, err := ac.Engine.ID(adminId).Cols("avatar").Update(&admin)
	return err != nil
}
