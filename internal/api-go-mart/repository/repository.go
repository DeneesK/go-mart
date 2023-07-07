package repository

import (
	"context"

	"github.com/DeneesK/go-mart/internal/api-go-mart/entities"
)

type Storage interface {
	Register(ctx context.Context, user *entities.User) error
	Login(ctx context.Context, user *entities.User) error

	GetUserOrders(ctx context.Context, login string) ([]byte, error)
	LoadOrder(ctx context.Context, login string, orderID string) error

	GetBalanceInfo(login string) ([]byte, error)
	Withdraw(login string, orderID string, sum float64) error
	GetWithdrawals(login string) ([]byte, error)
}
