package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
	"github.com/lovelyoyrmia/ecommerce/pkg/config"
	"github.com/lovelyoyrmia/ecommerce/pkg/fake"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	database := db.NewDatabase(context.Background(), c)
	store := db.NewStore(database.DB)

	seedProducts(store)
}

func seedProducts(store db.Store) {
	for i := 0; i < 20; i++ {
		product, err := store.CreateProduct(context.Background(), db.CreateProductParams{
			Pid:  uuid.NewString(),
			Name: fake.RandomString(),
			Description: pgtype.Text{
				String: fake.RandomString(),
				Valid:  true,
			},
			Category: pgtype.Int4{
				Int32: int32(fake.RandomInt(1, 4)),
				Valid: true,
			},
			Price: int32(fake.RandomPrice()),
			Stock: int32(fake.RandomStock()),
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Success Create Product : %v\n", product)
	}
}
