package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/slapshit-apocrypha/api/internal/web"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

type options struct {
	Addr       string `env:"SLAPSHIT_API_ADDR" envDefault:":8080"`
	WebhookURL string `env:"SLAPSHIT_API_WEBHOOK_URL"`
	Debug      bool   `env:"SLAPSHIT_API_DEBUG" envDefault:"false"`
}

func die(msg string, args ...any) {
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	fmt.Fprintf(os.Stderr, msg, args...)
}

func main() {
	envOpts := env.Options{
		RequiredIfNoDef: true,
	}
	var opts options

	if err := env.ParseWithOptions(&opts, envOpts); err != nil {
		die("error parsing environment options: %s\n", err.Error())
		return
	}

	logger := ctxlog.New(opts.Debug)
	ctx := ctxlog.WithLogger(context.Background(), logger)

	wc, err := webhook.NewFromURL(opts.WebhookURL)
	if err != nil {
		logger.Fatal("error creating webhook client", zap.Error(err))
	}

	webOpts := []web.Option{
		web.WithAddr(opts.Addr),
		web.WithWebhook(wc),
	}

	if opts.Debug {
		webOpts = append(webOpts, web.WithDebug())
	}

	srv, err := web.New(webOpts...)
	if err != nil {
		logger.Fatal("error creating web server", zap.Error(err))
	}

	if err := srv.Run(ctx); err != nil {
		logger.Fatal("error starting server", zap.Error(err))
	}
}
