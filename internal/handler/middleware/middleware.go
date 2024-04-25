package middleware

type Dependencies struct {
}

type Middleware struct {
}

func NewMiddleware(_ Dependencies) *Middleware {
	return &Middleware{}
}
