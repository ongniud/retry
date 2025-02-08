package retry

import (
	"math"
	"time"
)

type DelayFunc func(attempt uint32) time.Duration

func ConstDelay(delay time.Duration) DelayFunc {
	return func(attempt uint32) time.Duration {
		return delay
	}
}

func NoDelay() DelayFunc {
	return func(attempt uint32) time.Duration {
		return 0
	}
}

func ExponentialDelay(delay time.Duration, maxDelay time.Duration) DelayFunc {
	if delay <= 0 {
		return ConstDelay(delay)
	}
	if delay > maxDelay {
		maxDelay = delay
	}

	// after this attempt, we would exceed maxDelay
	maxAttempts := uint32(math.Floor(math.Log2(float64(maxDelay/delay)))) + 1
	if maxAttempts >= 64 {
		maxAttempts = 63
	}

	return func(attempt uint32) time.Duration {
		if attempt <= 0 {
			return delay
		} else if attempt > maxAttempts {
			return maxDelay
		}

		exp := int64(1) << uint(attempt-1)
		d := time.Duration(exp) * delay
		if d > maxDelay {
			return maxDelay
		}

		return d
	}
}
