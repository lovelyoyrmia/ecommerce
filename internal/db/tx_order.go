package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

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
			newOrder, err := store.CreateOrder(ctx, CreateOrderParams{
				Oid:         newOid,
				Uid:         req.Uid,
				TotalAmount: req.Quantity * req.Amount,
			})
			if err != nil {
				return err
			}
			return q.CreateOrderItems(ctx, CreateOrderItemsParams{
				Oid:      newOrder.Oid,
				Pid:      req.Pid,
				Quantity: req.Quantity,
				Amount:   req.Amount,
			})
		}

		err = q.CreateOrderItems(ctx, CreateOrderItemsParams{
			Oid:      order.Oid,
			Pid:      req.Pid,
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
	})
}
