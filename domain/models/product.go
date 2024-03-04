package models

type Product struct {
	Pid         string `json:"pid"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
	Stock       int32  `json:"stock"`
}

type ProductsParams struct {
	Limit    *int32
	Offset   *int32
	Category *string
}

type ProductDetailsParams struct {
	Pid string
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}
