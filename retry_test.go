package retry

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"
)

func TestRetry_DoCtxWithMaxRetries(t *testing.T) {
	err := Do(
		context.Background(),
		func(_ context.Context) error {
			return nil
		},
		WithMaxTries(6),
		WithDelayFunc(ConstDelay(1)),
	)
	_ = err
}

func TestRetry_DoCtxWithTimeout(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*150)

	start := time.Now()

	err := Do(
		ctx,
		func(_ context.Context) error {
			time.Sleep(time.Second * 5)
			e := errors.New("ss")
			//return e
			return MarkRetryable(e)
		},
		WithMaxTries(3),
		//WithDelayFunc(DelayConst(time.Second)),
	)

	cost := time.Since(start)

	log.Print("cost:", cost, "  err:", err)

	_ = err
}
