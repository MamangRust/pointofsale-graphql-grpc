package middlewares

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/resilience"
)

type ResilienceHttpMiddleware struct {
	LoadMonitor    *resilience.LoadMonitor
	CircuitBreaker *resilience.CircuitBreaker
	RequestLimiter *resilience.RequestLimiter
}

func NewResilienceHttpMiddleware(
	lm *resilience.LoadMonitor,
	cb *resilience.CircuitBreaker,
	rl *resilience.RequestLimiter,
) *ResilienceHttpMiddleware {
	return &ResilienceHttpMiddleware{
		LoadMonitor:    lm,
		CircuitBreaker: cb,
		RequestLimiter: rl,
	}
}

func (m *ResilienceHttpMiddleware) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m.LoadMonitor.RecordRequest()

			if m.CircuitBreaker.IsOpen() {
				http.Error(
					w,
					"Service temporarily unavailable due to high error rate. Please try again later.",
					http.StatusServiceUnavailable,
				)
				return
			}

			if !m.RequestLimiter.TryAcquire() {
				http.Error(
					w,
					"Too many concurrent requests. Please try again later.",
					http.StatusTooManyRequests,
				)
				return
			}

			defer m.RequestLimiter.Release()

			next.ServeHTTP(w, r)
		})
	}
}
