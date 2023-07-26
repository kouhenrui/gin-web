package util

import (
	"HelloGin/src/dto/reqDto"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"regexp"
)

/*
 * @MethodName ExistIn
 * @Description 判断参数是否存在
 * @Author khr
 * @Date 2023/4/14 8:52
 */

func ExistIn(param string, paths []string) bool {
	for _, v := range paths {
		if param == v {
			return true
		}
	}
	return false
}

/*
 * @MethodName FuzzyMatch
 * @Description 正则模糊匹配路径
 * @Author khr
 * @Date 2023/5/9 16:25
 */
func FuzzyMatch(param string, paths []string) bool {
	for _, y := range paths {
		if regexp.MustCompile(y).MatchString(param) {

			//fmt.Print("匹配道路进了")
			return true
		}

	}
	return false
}

var CaptchaStore base64Captcha.Store = RedisStore{}

/*
 * @MethodName CreateCaptcha
 * @Description 生成图片验证
 * @Author khr
 * @Date 2023/5/8 10:44
 */

func CreateCaptcha() (error, interface{}) {
	var newCaptcha reqDto.Captcha
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          80,                                     //高度
		Width:           240,                                    //宽度
		NoiseCount:      0,                                      //干扰数
		ShowLineOptions: 6,                                      //展示个数
		Length:          6,                                      //长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", //验证码随机字符串来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体
	}
	driver = driverString.ConvertFonts()
	//生成验证码
	c := base64Captcha.NewCaptcha(driver, CaptchaStore)
	id, content, err := c.Generate()
	newCaptcha.Id = id
	newCaptcha.Content = content
	if err != nil {
		fmt.Println("生成有错:", err)
		return err, nil
	}
	return nil, &newCaptcha

}

/*
 * @MethodName VerifyCaptcha
 * @Description 验证图片验证码
 * @Author khr
 * @Date 2023/5/8 10:45
 */

func VerifyCaptcha(capt reqDto.Captcha) bool {
	// id 验证码id
	// answer 需要校验的内容
	// clear 校验完是否清除
	if CaptchaStore.Verify(capt.Id, capt.Content, true) {

		fmt.Println("验证正确")
		return true
	} else {
		fmt.Println("验证cuowu ")
		return false
	}
	//return nil, store.Verify(capt.Id, capt.Content, true)
}
