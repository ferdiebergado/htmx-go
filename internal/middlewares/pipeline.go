package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreatePipeline(layers ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(layers) - 1; i >= 0; i-- {
			layer := layers[i]
			next = layer(next)
		}

		return next
	}
}
