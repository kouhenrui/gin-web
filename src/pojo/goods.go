package pojo

import (
	"gin-web/src/dto/reqDto"
	"gin-web/src/dto/resDto"
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	ShopId      int     `json:"shop_id" `                //商家id
	Name        string  `json:"name" gorm:"not null"`    //名称
	Price       float32 `json:"price" gorm:"not null"`   //价格
	Address     string  `json:"address" gorm:"not null"` //地址
	Fms         string  `json:"fms"`                     //快递服务（免运费，现货，48小时发货）
	Brand       string  `json:"brand"`                   //品牌
	Producer    string  `json:"producer"`                //产地
	PriceRange  string  `json:"price_range"`             //价格区间
	Material    string  `json:"material"`                //材料
	ArticleNo   string  `json:"article_no"`              //货号
	Appraise    string  `json:"appraise"`                //评价 关联评价表
	Question    int     `json:"question"`                //问题 关联问题表
	ShowPic     string  `json:"show_pic"`                //展示图片
	Rotation    string  `json:"rotation" `               //轮播
	Description string  `json:"description"`             //描述
	Num         int     `json:"num" gorm:"not null"`     //数量
	SellNum     int     `json:"sell_num"`                //已售数量
	Status      string  `json:"status"`                  //商品状态 上架 grouding，违规 Violation，下架 undercarriage，缺货 missing
}

func GoodsServiceImpl() Goods {
	return Goods{}
}

var (
	//admins=[]Admin{}
	goodsInfo    = resDto.GoodsInfo{}
	goods        = Goods{}
	resGoodsList = []resDto.GoodsList{} //要查询的字段
)

// 增加
func (g *Goods) GoodsAdd(add Goods) error {
	return db.Save(&add).Error

}

// 查名字
func (g *Goods) GoodsByName(name string) (error, *resDto.GoodsInfo) {
	err := db.Model(&g).Where("name = ?", name).Find(&goodsInfo).Error
	if err != nil {
		return err, nil
	}
	return nil, &goodsInfo
}

// 查id
func (g *Goods) GoodsById(id uint) (error, *resDto.GoodsInfo) {
	err := db.Model(&g).Where("id = ?", id).Find(&goodsInfo).Error
	if err != nil {
		return err, nil
	}
	return nil, &goodsInfo
}

// 列表，支持模糊查询
func (g *Goods) GoodsList(list reqDto.GoodsList) (error, *resDto.CommonList) {
	query := db.Model(&g)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	err := query.Limit(list.Take).Offset(list.Skip).Find(&resGoodsList).Count(&count).Error
	if err != nil {
		return err, nil
	}
	reslist.List = resGoodsList
	reslist.Count = uint(count)
	return nil, &reslist

}

// 删除
func (g *Goods) GoodsDel(id uint) error {
	g.ID = id
	err := db.Delete(&g).Error
	if err != nil {
		return err
	}
	return nil
}

// 修改
func (g *Goods) GoodsUpdate(update Goods) error {
	return db.Model(&g).Updates(&update).Error

}
