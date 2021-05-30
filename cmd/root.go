package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/andersnormal/pkg/debug"
	"github.com/andersnormal/pkg/opts"
	"github.com/andersnormal/pkg/server"
	"github.com/spf13/cobra"

	"github.com/katallaxie/puzzler/pkg/puzzler"
)

var root = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		s, _ := server.WithContext(ctx)

		d := debug.New(
			debug.WithPprof(),
			debug.WithStatusAddr(":8443"),
		)

		o := opts.NewDefaultOpts()
		p := puzzler.New(o)

		s.Listen(d, false)
		s.Listen(p, true)

		err := s.Wait()
		var e *server.ServerError
		if errors.As(err, &e) {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
