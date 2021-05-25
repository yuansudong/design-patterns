package proxy

type Service interface {
	HandleRequest(string, string) (int, string)
}

type Nginx struct {
	application *UserService
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application: &UserService{},
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	return n.application.HandleRequest(url, method)
}

type UserService struct {
}

func (a *UserService) HandleRequest(url, method string) (int, string) {
	if url == "/user" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}
