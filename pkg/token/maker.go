package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type Maker interface {
	GenerateToken(uid string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
