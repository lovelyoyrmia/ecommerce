package models

import "github.com/lovelyoyrmia/ecommerce/internal/db"


type Product struct {
	Pid         string `json:"pid"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
	Stock       int32  `json:"stock"`
}

type ProductsParams struct {
	Limit  int32
	Offset int32
}

type ProductDetailsParams struct {
	Pid string
}

type ProductsResponse struct {
	Products []db.GetProductsRow `json:"products"`
}
