package retry

type Options struct {
	maxTries  uint32
	delayFunc DelayFunc
}

func newRetryOptions(options ...Option) *Options {
	opts := &Options{
		delayFunc: NoDelay(),
	}
	for _, option := range options {
		option(opts)
	}
	return opts
}

type Option func(opts *Options)

func WithMaxTries(tries uint32) Option {
	return func(opts *Options) {
		opts.maxTries = tries
	}
}

func WithDelayFunc(df DelayFunc) Option {
	return func(opts *Options) {
		opts.delayFunc = df
	}
}
