package main

import (
	"fmt"

	"github.com/yuansudong/design-patterns/proxy"
)

func main() {
	nginx := proxy.NewNginxServer()
	fmt.Println(nginx.HandleRequest("/user", "GET"))
	fmt.Println(nginx.HandleRequest("/user", "POST"))
}
