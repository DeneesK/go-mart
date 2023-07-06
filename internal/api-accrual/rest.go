package apiaccrual

import (
	"net/http"

	"github.com/go-chi/chi"
)

const (
	contentType     = "application/json"
	contentTypeText = "text/plain"
)

type Storage interface {
}

func Routers(db Storage) chi.Router {
	r := chi.NewRouter()
	r.Get("/api/orders/{number}", OrderInfo(db))
	return r
}

func OrderInfo(db Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		_ = chi.URLParam(req, "number")
		res.WriteHeader(http.StatusOK)
		res.Header().Add("Content-Type", contentTypeText)
	}
}
