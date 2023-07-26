package pojo

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"` //名称
	//Price       float32 `json:"price" gorm:"not null"`    //价格
	Address string `json:"address" gorm:"not null"` //地址
	Fms     string `json:"fms"`                     //快递服务（免运费，现货，48小时发货）
	//Brand       string  `json:"brand"`                    //品牌
	//Producer    string  `json:"producer"`                 //产地
	//PriceRange  string  `json:"price_range"`              //价格区间
	//Material    string  `json:"material"`                 //材料
	//ArticleNo   string  `json:"article_no"`               //货号
	//Appraise    string  `json:"appraise"`                 //评价 关联评价表
	//Question    int     `json:"question"`                 //问题 关联问题表
	//ShowPic     string  `json:"show_pic" gorm:"not null"` //展示图片
	//Rotation    string  `json:"rotation" `                //轮播
	//Description string  `json:"description"`              //描述
	//Num         int     `json:"num" gorm:"not null"`      //数量
	SellNum int    `json:"sell_num"`               //已售数量
	Status  string `json:"status" gorm:"not null"` //商点状态 上架，违规，下架，缺货
}
