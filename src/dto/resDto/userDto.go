package resDto

/**
* @program: work_space
*
* @description:返回参数格式化
*
* @author: khr
*
* @create: 2023-02-01 14:15
**/

type UserList struct {
	Id      uint   `json:"id"`
	Name    string `json:"name" `
	Account string `json:"account"`
	Role    int    `json:"role"`
}

// 详情数据
type UserInfo struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Role     int    `json:"role"`
	RoleName string `json:"role_name"`
}
