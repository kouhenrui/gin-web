package comDto

/**
 * @ClassName common
 * @Description TODO
 * @Author khr
 * @Date 2023/5/8 13:37
 * @Version 1.0
 */

type TokenClaims struct {
	Id       uint   `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
	Account  string `json:"account"`
	RoleName string `json:"role_name""`
}
