package middlewares

import (
	"github.com/lovelyoyrmia/ecommerce/internal/db"
	"github.com/lovelyoyrmia/ecommerce/pkg/config"
	"github.com/lovelyoyrmia/ecommerce/pkg/token"
)

type Middlewares struct {
	config config.Config
	maker  token.Maker
	store  db.Store
}

func NewMiddlewares(config config.Config, maker token.Maker, store db.Store) *Middlewares {
	return &Middlewares{
		config: config,
		maker:  maker,
		store:  store,
	}
}
