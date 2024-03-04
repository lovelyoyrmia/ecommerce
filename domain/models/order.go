package models

type CartParams struct {
	Uid      string
	Pid      string `json:"pid"`
	Quantity int32  `json:"quantity"`
	Amount   int32  `json:"amount"`
}

type CartProductParams struct {
	Oid string
	Uid string
	Pid string
}

type CartsParams struct {
	Uid string
}

type CartProduct struct {
	Pid          string `json:"pid"`
	Name         string `json:"name"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Price        int64  `json:"price"`
	Stock        int32  `json:"stock"`
	Quantity     int32  `json:"quantity"`
	Amount       int64  `json:"amount"`
}

type CartProducts struct {
	Oid      string
	Products []CartProduct
}

type Orders struct {
	Oid         string `json:"oid"`
	OrderStatus string `json:"order_status"`
	OrderedAt   string `json:"ordered_at"`
}

func (o *Orders) ExtractOrderStatus(orderStatus int32) Orders {
	var order string
	switch orderStatus {
	case 1:
		order = "pending"
	case 2:
		order = "processed"
	case 3:
		order = "completed"
	case 4:
		order = "canceled"
	default:
		order = "added"
	}
	return Orders{
		Oid:         o.Oid,
		OrderStatus: order,
		OrderedAt:   o.OrderedAt,
	}
}
