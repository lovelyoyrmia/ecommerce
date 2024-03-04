package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrOutOfStock = errors.New("out of stock")

type CreateCartTx struct {
	Uid      string
	Pid      string
	Amount   int32
	Quantity int32
}

func (store *SQLStore) CreateCartUserTx(ctx context.Context, req CreateCartTx) error {
	return store.ExecTx(ctx, func(q *Queries) error {
		order, err := q.GetCartUser(ctx, req.Uid)

		if errors.Is(err, ErrRecordNotFound) {
			newOid := uuid.NewString()
			product, err := q.GetProductDetails(ctx, req.Pid)
			if err != nil {
				return err
			}
			if product.Stock <= req.Quantity {
				return ErrOutOfStock
			}
			newOrder, err := q.CreateOrder(ctx, CreateOrderParams{
				Oid:         newOid,
				Uid:         req.Uid,
				TotalAmount: req.Quantity * req.Amount,
			})
			if err != nil {
				return err
			}
			return q.CreateOrderItems(ctx, CreateOrderItemsParams{
				Oid:      newOrder.Oid,
				Pid:      product.Pid,
				Quantity: req.Quantity,
				Amount:   req.Amount * req.Quantity,
			})
		}

		product, err := q.GetProductDetails(ctx, req.Pid)
		if err != nil {
			return err
		}
		if product.Stock < req.Quantity {
			return ErrOutOfStock
		}

		orderItem, err := q.GetCartProductDetail(ctx, GetCartProductDetailParams{
			Oid: order.Oid,
			Pid: product.Pid,
		})

		if product.Stock < (orderItem.Quantity + req.Quantity) {
			return ErrOutOfStock
		}

		if errors.Is(err, ErrRecordNotFound) {
			err = q.CreateOrderItems(ctx, CreateOrderItemsParams{
				Oid:      order.Oid,
				Pid:      product.Pid,
				Quantity: req.Quantity,
				Amount:   req.Amount,
			})
			if err != nil {
				return err
			}
			return q.UpdateCart(ctx, UpdateCartParams{
				TotalAmount: order.TotalAmount + (req.Quantity * req.Amount),
				Oid:         order.Oid,
			})
		}
		if err != nil {
			return err
		}

		err = q.UpdateCartProductDetail(ctx, UpdateCartProductDetailParams{
			Oid:      order.Oid,
			Pid:      product.Pid,
			Quantity: orderItem.Quantity + req.Quantity,
			Amount:   orderItem.Amount + (req.Amount * req.Quantity),
		})
		if err != nil {
			return err
		}
		return q.UpdateCart(ctx, UpdateCartParams{
			TotalAmount: order.TotalAmount + (orderItem.Amount + (req.Amount * req.Quantity)),
			Oid:         order.Oid,
		})
	})
}
