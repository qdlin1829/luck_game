package model

type OrderSerach struct {
	OrderSn       string    `json:"order_sn" form:"order_sn"`
	OrderStatus   int       `json:"order_status" form:"order_status"`
	Page          int       `json:"page" form:"page"`
	PageSize      int       `json:"page_size" form:"page_size"`
	UserId        int       `json:"user_id" form:"user_id"`
}
