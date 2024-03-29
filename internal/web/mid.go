package web

import (
	"net/http"

	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func injectLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := ctxlog.WithLogger(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				ctx := ctxlog.WithOptions(r.Context(), zap.AddStacktrace(zap.ErrorLevel))
				ctxlog.Error(ctx, "PANIC", zap.Any("val", rvr))

				respond(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxlog.Debug(r.Context(), "request",
			zap.String("method", r.Method),
			zap.String("url", r.RequestURI),
			zap.String("protocol", r.Proto),
		)

		next.ServeHTTP(w, r)
	})
}
