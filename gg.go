package gg

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Runner is a thing that runs with a context and returns an error.
type Runner interface {
	Run(context.Context) error
}

// Run the runners with the provided context.
// They each run in their own goroutine.
func Run(ctx context.Context, rs ...Runner) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, r := range rs {
		eg.Go(func(runner Runner) func() error {
			return func() error {
				return r.Run(ctx)
			}
		}(r))
	}
	return eg.Wait()
}
