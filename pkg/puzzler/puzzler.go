package puzzler

import (
	"context"
	"net/http"

	"github.com/katallaxie/puzzler/pkg/handler"

	o "github.com/andersnormal/pkg/opts"
	"github.com/andersnormal/pkg/server"
)

type puzzler struct {
	opts o.Opts
}

func New(opts o.Opts) server.Listener {
	return &puzzler{
		opts: opts,
	}
}

func (s *puzzler) Start(ctx context.Context, ready server.ReadyFunc) func() error {
	return func() error {
		http.HandleFunc("/", handler.Handler)

		if err := http.ListenAndServe(":12345", nil); err != nil {
			return err
		}

		return nil
	}
}
