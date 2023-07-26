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
type AdminList struct {
	Id      uint   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Account string `json:"account,omitempty"`
	Role    int    `json:"role,omitempty"`
}

// 详情数据
type AdminInfo struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Role     int    `json:"role"`
	RoleName string `json:"role_name"`
}

//type AdminInformation struct {
//
//}
