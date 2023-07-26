package reqDto

type TokenDate struct {
	Id          uint   `json:"id"`
	Name        string `json:"name,omitempty"`
	Account     string `json:"account,omitempty"`
	Salt        string `json:"salt,omitempty"`
	Role        int    `json:"role,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
type LoginRedisDate struct {
	Token   string `json:"token"`
	Exptime string `json:"exptime"`
}
