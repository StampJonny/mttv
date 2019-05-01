package context

type Context interface{}

type contextImpl struct {
	Context
}

func New() Context {
	return &contextImpl{}
}
