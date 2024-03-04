package addressparser

type Options struct {
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return options
}
