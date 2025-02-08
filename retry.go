package retry

import (
	"context"
	"time"
)

func Do(ctx context.Context, fn func(context.Context) error, retryOptions ...Option) error {
	opts := newRetryOptions(retryOptions...)
	attempts := uint32(0)
	for {
		err := fn(ctx)
		if err == nil {
			return nil
		}
		if r, ok := err.(IRetryable); !ok || !r.Retryable() {
			return err
		}
		attempts++
		if opts.maxTries != 0 && attempts >= opts.maxTries {
			return err
		}
		t := time.NewTimer(opts.delayFunc(attempts))
		select {
		case <-t.C:
		case <-ctx.Done():
			if !t.Stop() { // once context cancelled, kill the timer if it hasn't fired, and return the last error we got
				<-t.C
			}
			return ctx.Err()
		}
	}
}
