package web

import "github.com/diamondburned/arikawa/v3/api/webhook"

// Option configures a server.
type Option func(*Server)

// WithAddr sets a server's address.
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithWebhook sets a server's webhook URL.
func WithWebhook(hook *webhook.Client) Option {
	return func(s *Server) {
		s.webhookClient = hook
	}
}
