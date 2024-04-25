package http

type opt struct {
	queryParser   bool
	bodyParser    bool
	bodyValidator bool
}

var defaultOtp = opt{
	queryParser:   false,
	bodyParser:    false,
	bodyValidator: false,
}

func newOpt(opts ...OptFunc) opt {
	out := defaultOtp
	for _, o := range opts {
		o(&out)
	}

	return out
}

type OptFunc func(*opt)

func WithQueryParser() OptFunc {
	return func(o *opt) {
		o.queryParser = true
	}
}

func WithBodyParser() OptFunc {
	return func(o *opt) {
		o.bodyParser = true
	}
}

func WithBodyValidator() OptFunc {
	return func(o *opt) {
		o.bodyValidator = true
	}
}
