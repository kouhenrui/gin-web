package resDto

/**
* @program: work_space
*
* @description:
*
* @author: khr
*
* @create: 2023-02-07 15:35
**/
type CommonList struct {
	Count uint        `json:"count,omitempty"`
	List  interface{} `json:"list,omitempty"`
}

type TokenAndExp struct {
	Token   string `json:"token,omitempty"`
	Exptime string `json:"exptime"`
}
