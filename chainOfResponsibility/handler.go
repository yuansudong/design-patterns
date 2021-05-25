package chainOfResponsibility

// Handler
type Handler interface {
	SetNext(Handler)
	Execute(IContext) error
}
