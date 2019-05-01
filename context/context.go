package context

type Context interface {
	IsValid() bool
}

type contextImpl struct {
	Context
}

func (c *contextImpl) IsValid() bool {
	return true
}

func New() Context {
	return &contextImpl{}
}
