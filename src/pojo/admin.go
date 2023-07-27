package pojo

import (
	"gin-web/src/dto/reqDto"
	"gin-web/src/dto/resDto"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Password    string `json:"password" `
	Salt        string `json:"salt"`
	Account     string `json:"account" gorm:"unique:true"`
	AccessToken string `json:"access_token"`
	Revoke      bool   `json:"revoke" gorm:"default:false"`
	Role        int    `json:"role"`
}

func AdminServiceImpl() Admin {
	return Admin{}
}

var (
	//admins=[]Admin{}
	adminInfo    = resDto.AdminInfo{}
	admin        = Admin{}
	resAdminList = []resDto.AdminList{} //要查询的字段
)

// 分页,模糊查询用户
func (a *Admin) AdminList(list reqDto.AdminList) (*resDto.CommonList, error) {
	query := db.Model(&a)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	err := query.Limit(list.Take).Offset(int(list.Skip)).Find(&resAdminList).Count(&count).Error
	reslist.Count = uint(count)
	reslist.List = resAdminList
	if err != nil {
		return nil, err
	}
	return &reslist, nil
}

// 查询账号
func (a *Admin) CheckByAccount(account string) (*Admin, error) {
	//a.Account = account
	a.Account = account
	err := db.First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 查询名称
func (a *Admin) CheckByName(name string) (*Admin, error) {
	a.Name = name
	err := db.First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 详情数据
func (a *Admin) AdminInfo(id int, name string) (*resDto.AdminInfo, error) {
	adminInfo.Name = name
	a.ID = uint(id)
	a.Name = name
	err := db.Model(&a).Select("admin.name, admin.account,admin.role,r.name as role_name").Joins("left join rule as r on r.id = admin.role").Scan(&adminInfo).Error
	if err != nil {
		return nil, err
	}
	return &adminInfo, nil
}

// 更新token数据
func (a *Admin) UpdateToken(access_token string, id uint) error {
	a.ID = id
	err := db.Model(&a).Update("access_token", access_token).Error
	//if res.RowsAffected != 1 {
	//	return false
	//}
	if err != nil {
		return err
	}
	return nil
}

// 增加用户
func (a *Admin) AddAdmin(admins Admin) error {
	err := db.Create(&admins).Error
	if err != nil {
		return err
	}
	return err
}
