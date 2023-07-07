package orders

import (
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/handlers"
	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func Orders(r chi.Router, handler *handlers.Handler) {
	r.Group(func(router chi.Router) {
		r.Use(middleware.JWTAuth)
		router.Post("/orders", handler.LoadOrders)
		router.Get("/orders", handler.GetOrders)
	})
}
