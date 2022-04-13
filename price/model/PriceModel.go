package model

import "time"

type Price struct {
	Channel    string    `json:"channel"`
	Sku        string    `json:"sku"`
	Price      int       `json:"price"`
	SkuName    string    `json:"skuName"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}
