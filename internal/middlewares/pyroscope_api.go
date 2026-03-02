package middlewares

import (
	"context"
	"net/http"

	"github.com/grafana/pyroscope-go"
)

func PyroscopeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pyroscope.TagWrapper(
			r.Context(),
			pyroscope.Labels(
				"endpoint", r.URL.Path,
				"method", r.Method,
			),
			func(ctx context.Context) {
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			},
		)
	})
}
