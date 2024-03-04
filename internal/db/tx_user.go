package db

import (
	"context"
)

func (store *SQLStore) CreateUserTx(ctx context.Context, userParams CreateUserParams) (User, error) {
	var userRes User
	err := store.ExecTx(ctx, func(q *Queries) error {
		user, err := q.CreateUser(ctx, userParams)
		if err != nil {
			return err
		}
		userRes = user
		return nil
	})
	if err != nil {
		return User{}, err
	}
	return userRes, nil
}
