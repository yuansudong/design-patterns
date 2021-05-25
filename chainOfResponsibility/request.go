package chainOfResponsibility

type Request struct {
	tok string
}

func NewRequest(t string) *Request {

	return &Request{
		tok: t,
	}
}

func (r *Request) GetToken() string {
	return r.tok
}
