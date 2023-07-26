package resDto

type GoodsInfo struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name" `       //名称
	Price       float32 `json:"price" `      //价格
	Address     string  `json:"address"`     //地址
	Fms         string  `json:"fms"`         //快递服务（免运费，现货，48小时发货）
	Brand       string  `json:"brand"`       //品牌
	Producer    string  `json:"producer"`    //产地
	PriceRange  string  `json:"price_range"` //价格区间
	Material    string  `json:"material"`    //材料
	ArticleNo   string  `json:"article_no"`  //货号
	Appraise    string  `json:"appraise"`    //评价 关联评价表
	Question    int     `json:"question"`    //问题 关联问题表
	ShowPic     string  `json:"show_pic"`    //展示图片
	Rotation    string  `json:"rotation" `   //轮播
	Description string  `json:"description"` //描述
	Num         int     `json:"num"`         //数量
	SellNum     int     `json:"sell_num"`    //已售数量
	Status      string  `json:"status"`      //商品状态 上架 grouding，违规 Violation，下架 undercarriage，缺货 missing
}

type GoodsList struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name" `       //名称
	Price       float32 `json:"price" `      //价格
	Address     string  `json:"address"`     //地址
	Fms         string  `json:"fms"`         //快递服务（免运费，现货，48小时发货）
	Brand       string  `json:"brand"`       //品牌
	Producer    string  `json:"producer"`    //产地
	PriceRange  string  `json:"price_range"` //价格区间
	Material    string  `json:"material"`    //材料
	ArticleNo   string  `json:"article_no"`  //货号
	Appraise    string  `json:"appraise"`    //评价 关联评价表
	Question    int     `json:"question"`    //问题 关联问题表
	ShowPic     string  `json:"show_pic"`    //展示图片
	Rotation    string  `json:"rotation" `   //轮播
	Description string  `json:"description"` //描述
	Num         int     `json:"num"`         //数量
	SellNum     int     `json:"sell_num"`    //已售数量
	Status      string  `json:"status"`      //商品状态 上架 grouding，违规 Violation，下架 undercarriage，缺货 missing
}
