package infra

import (
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

type Router struct {
	c *controller.AccountItemController
}

func NewRouter(c *controller.AccountItemController) Router { return Router{c} }

func (r Router) Run() {
	http.HandleFunc("/", r.c.Get)
	http.ListenAndServe("localhost:8080", nil)
}
