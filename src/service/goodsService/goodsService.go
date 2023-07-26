package goodsService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/pojo"
)

// var goods=pojo.Goods{}
var goodsServiceImpl = pojo.GoodsServiceImpl()

var err error

// 增加
func GoodsAdd(add reqDto.GoodsAdd) error {
	goods := pojo.Goods{
		Name:        add.Name,
		Price:       add.Price,
		Address:     add.Address,
		Fms:         add.Fms,
		Brand:       add.Brand,
		Producer:    add.Producer,
		PriceRange:  add.PriceRange,
		Material:    add.Material,
		ArticleNo:   add.ArticleNo,
		Appraise:    add.Appraise,
		Question:    add.Question,
		ShowPic:     add.ShowPic,
		Rotation:    add.Rotation,
		Description: add.Description,
		Num:         add.Num,
		SellNum:     add.SellNum,
		Status:      add.Status}
	return goodsServiceImpl.GoodsAdd(goods)
}

// 通过id查询
func GoodsFindById(id int) (error, *resDto.GoodsInfo) {
	var info = &resDto.GoodsInfo{}
	err, info = goodsServiceImpl.GoodsById(uint(id))
	if err != nil {
		return err, nil
	}
	return nil, info
}

// 通过名称查询
func GoodsFindByName(name string) (error, *resDto.GoodsInfo) {
	var info = &resDto.GoodsInfo{}
	err, info = goodsServiceImpl.GoodsByName(name)
	if err != nil {
		return err, nil
	}
	return nil, info
}

// 修改商品
func GoodsUpdate(update reqDto.GoodsUpdate) error {
	goods := pojo.Goods{
		Name:        update.Name,
		Price:       update.Price,
		Address:     update.Address,
		Fms:         update.Fms,
		Brand:       update.Brand,
		Producer:    update.Producer,
		PriceRange:  update.PriceRange,
		Material:    update.Material,
		ArticleNo:   update.ArticleNo,
		Appraise:    update.Appraise,
		Question:    update.Question,
		ShowPic:     update.ShowPic,
		Rotation:    update.Rotation,
		Description: update.Description,
		Num:         update.Num,
		SellNum:     update.SellNum,
		Status:      update.Status}
	goods.ID = update.Id
	return goodsServiceImpl.GoodsUpdate(goods)
}

// 列表
func GoodsList(list reqDto.GoodsList) (error, *resDto.CommonList) {
	var goodsList = &resDto.CommonList{}
	err, goodsList = goodsServiceImpl.GoodsList(list)
	if err != nil {
		return err, nil
	}
	return nil, goodsList
}
