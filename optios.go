package addressparser

type Options struct {
	TextFilter []string
	NameMaxLen int
}

type Option func(o *Options)

func TextFilter(f []string) Option {
	return func(o *Options) {
		o.TextFilter = f
	}
}

func NameMaxLen(l int) Option {
	return func(o *Options) {
		o.NameMaxLen = l
	}
}

func NewOptions(opts ...Option) Options {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func NewAddressParser(opts ...Option) *AddressParser {
	options := NewOptions(opts...)

	return &AddressParser{
		opts: options,
	}
}
