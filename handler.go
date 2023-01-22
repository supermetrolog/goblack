package goblack

type Handler interface {
	Handler(c Context) (Response, error)
}

type Middleware interface {
	Handler(c Context, next Handler) (Response, error)
}
