package model

type PriceModel struct {
	Id      int
	SkuName string
	Sku     string
}

func (PriceModel) TableName() string {
	return "t_price"
}
