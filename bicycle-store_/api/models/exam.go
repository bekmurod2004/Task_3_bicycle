package models

type Promo struct {
	Promo_id    int    `json:"promo_id"`
	PromoName   string `json:"promo_name"`
	IsPercent   bool   `json:"is_percent"`
	Discount    int    `json:"discount"`
	Limit_Price int    `json:"order_limit_price"`
}

type PromoCreate struct {
	PromoName   string `json:"promo_name"`
	IsPercent   bool   `json:"is_percent"`
	Discount    int    `json:"discount"`
	Limit_Price int    `json:"order_limit_price"`
}
type PromoPrimaryKey struct {
	Promo_id int `json:"promo_id"`
}

type Query struct {
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
	Search string `json:"search"`
}
