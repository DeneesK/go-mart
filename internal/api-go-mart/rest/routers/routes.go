package routers

import (
	"github.com/DeneesK/go-mart/internal/api-go-mart/repository"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/handlers"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/routers/auth"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/routers/balance"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/routers/orders"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/routers/withdrawals"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func GetRoutes(r *chi.Mux, db repository.Storage, log *zap.SugaredLogger) {
	handler := handlers.New(db, log)
	r.Route("/api/user/", func(r chi.Router) {
		auth.AuthRouters(r, handler)
		orders.Orders(r, handler)
		balance.Balance(r, handler)
		withdrawals.Withdrawals(r, handler)
	})
}
