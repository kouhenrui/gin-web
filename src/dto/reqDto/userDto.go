package reqDto

type UserLogin struct {
	Account  string `json:"account" `
	Name     string `json:"name"  `
	Password string `json:"password" binding:"required" `
	Method   string `json:"method" binding:"required" gorm:"default:false;one of account,name"`
	Revoke   bool   `json:"revoke" validate:"required"`
}

// 图形验证
type Captcha struct {
	Id      string `json:"id,omitempty" `
	Content string `json:"content,omitempty" `
}
type TextMessage struct {
	Code string `json:"code,omitempty"`
}
type NewUserLogin struct {
	Phone    string      `json:"phone,omitempty" binding:"required" `
	Password string      `json:"password,omitempty" `
	Captcha  Captcha     `json:"captcha"`
	Message  TextMessage `json:"message"`
	Method   string      `json:"method" binding:"required" gorm:"one of captcha,message"`
	Revoke   bool        `json:"revoke,omitempty" validate:"required"`
}

type SignUser struct {
	Password string  `json:"password,omitempty"  binding:"required" validate:"required"`
	Phone    string  `json:"phone" binding:"required" validate:"omitempty"`
	Captcha  Captcha `json:"captcha" binding:"required"`
}

type UpdateUser struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Account  string `json:"account"`
}
type AddUser struct {
	Name     string `json:"name,omitempty" `
	Password string `json:"password,omitempty"  binding:"required" validate:"required"`
	Account  string `json:"account,omitempty"  binding:"required" validate:"omitempty"`
	Phone    string `json:"phone" binding:"required" validate:"omitempty"`
	Salt     string `json:"salt,omitempty"`
	Role     int    `json:"role,omitempty"`
}
type UserList struct {
	Take int    `json:"take,omitempty" binding:"required"`
	Skip uint   `json:"skip,omitempty"`
	Name string `json:"name,omitempty"`
}
