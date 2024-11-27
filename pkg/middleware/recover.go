package middleware

import (
	"context"
	"net/http"
)

type audiTrail struct {
	Method    string
	URL       string
	UserAgent string
}

func Recover() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			auditrail := audiTrail{
				Method:    r.Method,
				URL:       r.URL.Path,
				UserAgent: r.UserAgent(),
			}

			ctx := context.WithValue(r.Context(), "auditrail", auditrail)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
