package pojo

import (
	"errors"
	"gin-web/src/dto/reqDto"
	"gin-web/src/dto/resDto"
	"gin-web/src/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" gorm:"default:隔壁老王"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	Account     string `json:"account" tag:"unique"`
	Phone       string `json:"phone"`
	AccessToken string `json:"access_token"`
	Revoke      bool   `json:"revoke" gorm:"default:false"`
	Role        int    `json:"role" gorm:"default:5;type:int"`
}

func UserServiceImpl() User {
	return User{}
}

var (
	userList     = &[]User{} //多个user返回
	user         = User{}
	resUsersList = []resDto.UserList{} //要查询的字段
)

func (u *User) CheckPhone(phone string) (error, *User) {
	u.Phone = phone
	err := db.First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理记录不存在的情况
		return errors.New(util.SQL_NOT_EXIT_ERROR), u
	} else {
		// 处理其他错误
		if err != nil {
			return err, u
		}
		return nil, u
	}
	//if err != nil {
	//	return err, u
	//}
	//return nil, u
}

// 分页,模糊查询用户
func (u *User) UserList(list reqDto.UserList) *resDto.CommonList {
	query := db.Model(user)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&resUsersList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = resUsersList
	return &reslist
}

// 查询账号
func (u *User) CheckByAccount(account string) (error, *User) {
	//var userInfo resDto.UserInfo
	//userInfo.Account = account
	u.Account = account
	err := db.First(&u).Error
	//fmt.Println(userInfo)
	if err != nil {
		return err, u
	}
	return nil, u
}

// 查询名称
func (u *User) CheckByName(name string) (error, *User) {
	//var userInfo resDto.UserInfo
	//userInfo.Name = name
	u.Name = name
	err := db.First(&u).Error
	if err != nil {
		return err, u
	}
	return nil, u
}

// 更新token数据
func (u *User) UpdateToken(access_token string, id uint) error {

	//u.ID = id
	//fmt.Println(u, "打印实体id")
	err := db.Model(&u).Update("access_token", access_token).Where("id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
		// 处理记录不存在的情况
	} else {
		// 处理其他错误
		return err
	}
	//fmt.Println("打印sql错误", err)

}

// 增加用户
func (u *User) AddUser(user User) error {
	return db.Create(&user).Error
}

//根据手机号查询
