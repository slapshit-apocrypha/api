package web

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/go-chi/chi/v5"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

const (
	defaultPort = ":8080"
)

// Server is an HTTP server responsible for serving slapshit.net's API.
type Server struct {
	addr          string
	webhookClient *webhook.Client
}

// New creates a new server.
func New(opts ...Option) (*Server, error) {
	s := &Server{}

	for _, o := range opts {
		o(s)
	}

	if s.addr == "" {
		s.addr = defaultPort
	}

	if s.webhookClient == nil {
		return nil, fmt.Errorf("web: missing webhook client")
	}

	return s, nil
}

// Run registers routes and starts a server.
func (s *Server) Run(ctx context.Context) error {
	r := chi.NewMux()

	r.Use(recoverer)
	logger := ctxlog.FromContext(ctx)
	r.Use(requestLogger(logger))

	r.Post("/dj/application", s.postSendWebhook)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})

	srv := &http.Server{
		Addr:        s.addr,
		Handler:     r,
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			ctxlog.Error(ctx, "error shutting down server", zap.Error(err))
			return
		}
	}()

	ctxlog.Info(ctx, "starting server", zap.String("addr", s.addr))
	return srv.ListenAndServe()
}
