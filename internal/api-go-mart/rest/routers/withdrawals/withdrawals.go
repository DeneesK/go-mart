package withdrawals

import (
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/handlers"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func Withdrawals(r chi.Router, handler *handlers.Handler) {
	r.Group(func(router chi.Router) {
		r.Use(middleware.JWTAuth)
		router.Get("/withdrawals", handler.Withdrawals)
	})
}
